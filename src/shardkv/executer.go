package shardkv

import (
	"6.5840/raft"
	"6.5840/shardctrler"
)

func (kv *ShardKV) execute() {
	for kv.killed() == false {
		msg := <-kv.applyCh
		if msg.CommandValid {
			err := Err("")
			if _, ok := msg.Command.(Op); ok {
				err = kv.applyOp(msg)
			} else if _, ok := msg.Command.(shardctrler.Config); ok {
				kv.applyConfig(msg)
			} else if _, ok := msg.Command.(PushShardArgs); ok {
				err = kv.applyPushShard(msg)
			} else if _, ok := msg.Command.(DeleteShardArgs); ok {
				err = kv.applyDeleteShard(msg)
			}

			kv.replyMsg(msg, err)
			kv.snapshot(msg)
		} else {
			kv.mu.Lock()
			snapshot := msg.Snapshot
			kv.readSnapshot(snapshot)
			Debug(dSnap, "G%d/S%d installs snapshot index=%d term=%d bytes=%d config=%d keys=%d", kv.gid, kv.me, msg.SnapshotIndex, msg.SnapshotTerm, len(snapshot), kv.config.Num, len(kv.data))
			kv.mu.Unlock()
		}
	}
}

func (kv *ShardKV) applyOp(msg raft.ApplyMsg) Err {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	op := msg.Command.(Op)
	if op.IsGet {
		clientId := op.ClientId
		flag1, flag2 := kv.checkShard(op.Key)
		if !flag1 || !flag2 {
			return ErrSHARDNOTREADY
		}
		if op.SeqNo > kv.table[clientId] {
			kv.table[clientId] = op.SeqNo
		}
		if msg.IsLeader && kv.table[clientId] == op.SeqNo {
			Debug(dKVOp, "G%d/S%d applies GET index=%d client=%d seq=%d key=%q config=%d", kv.gid, kv.me, msg.CommandIndex, op.ClientId, op.SeqNo, op.Key, kv.config.Num)
			return OK
		}

	} else if op.IsPutAppend {
		key := op.Key
		value := op.Value
		clientId := op.ClientId
		flag1, flag2 := kv.checkShard(op.Key)
		if !flag1 || !flag2 {
			return ErrSHARDNOTREADY
		}
		if kv.isRepeated(op.ClientId, op.SeqNo) {
			Debug(dDedup, "G%d/S%d skips duplicate %s index=%d client=%d seq=%d", kv.gid, kv.me, op.PutOrAppend, msg.CommandIndex, op.ClientId, op.SeqNo)
			return ErrRepeated
		}
		if op.PutOrAppend == "Put" {
			kv.data[key] = value
		} else {
			kv.data[key] = kv.data[key] + value
		}
		kv.table[clientId] = op.SeqNo
		Debug(dKVOp, "G%d/S%d applies %s index=%d client=%d seq=%d key=%q value=%q config=%d", kv.gid, kv.me, op.PutOrAppend, msg.CommandIndex, op.ClientId, op.SeqNo, op.Key, summarizeValue(kv.data[key]), kv.config.Num)
		return OK
	}
	return ""
}

func (kv *ShardKV) applyConfig(msg raft.ApplyMsg) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	config := msg.Command.(shardctrler.Config)
	if config.Num == kv.config.Num+1 && kv.readyForNewConfig() {
		oldNum := kv.config.Num
		kv.lastConfig = kv.config
		kv.config = config
		kv.updateShardsState()
		Debug(dConfig, "G%d/S%d applies config %d -> %d shards=%v %s", kv.gid, kv.me, oldNum, kv.config.Num, kv.config.Shards, shardStateSummary(kv.shardsState))
	}
}

func (kv *ShardKV) applyPushShard(msg raft.ApplyMsg) Err {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	args := msg.Command.(PushShardArgs)
	if args.Num < kv.config.Num {
		return OK
	}

	if args.Num > kv.config.Num {
		return ErrSHARDNOTREADY
	}

	pushed := true
	for _, shard := range args.Shards {
		if kv.shardsState[shard] != WORKING {
			pushed = false
		}
	}
	if pushed {
		Debug(dDedup, "G%d/S%d already installed shards=%v config=%d", kv.gid, kv.me, args.Shards, args.Num)
		return OK

	}
	for _, shard := range args.Shards {
		kv.shardsState[shard] = WORKING
	}
	for k, v := range args.Data {
		kv.data[k] = v
	}
	for k, v := range args.Table {
		if v > kv.table[k] {
			kv.table[k] = v
		}
	}
	Debug(dMigrate, "G%d/S%d installs shards=%v config=%d keys=%d clients=%d %s", kv.gid, kv.me, args.Shards, args.Num, len(args.Data), len(args.Table), shardStateSummary(kv.shardsState))
	return OK
}

func (kv *ShardKV) applyDeleteShard(msg raft.ApplyMsg) Err {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	args := msg.Command.(DeleteShardArgs)
	if args.Num < kv.config.Num {
		return OK
	}

	if args.Num > kv.config.Num {
		return ErrSHARDNOTREADY
	}

	if kv.config.Num == args.Num {
		deleted := false
		for _, shard := range args.Shards {
			if kv.shardsState[shard] == WORKING {
				deleted = true
			}
		}
		if deleted {
			Debug(dDedup, "G%d/S%d already garbage-collected shards=%v config=%d", kv.gid, kv.me, args.Shards, args.Num)
			return OK
		}
	}

	for _, shard := range args.Shards {
		kv.shardsState[shard] = WORKING
	}
	for _, k := range args.Keys {
		delete(kv.data, k)
	}
	Debug(dMigrate, "G%d/S%d garbage-collects shards=%v config=%d keys=%d %s", kv.gid, kv.me, args.Shards, args.Num, len(args.Keys), shardStateSummary(kv.shardsState))
	return OK

}

func (kv *ShardKV) replyMsg(msg raft.ApplyMsg, str Err) {
	kv.mu.Lock()
	ch, ok := kv.waitCh[msg.CommandIndex]
	kv.mu.Unlock()
	if ok {
		if msg.IsLeader {
			select {
			case ch <- str:
			default:
			}
		} else {
			select {
			case ch <- ErrNotApplied:
			default:
			}
		}
	}

}

func (kv *ShardKV) snapshot(msg raft.ApplyMsg) {
	kv.mu.Lock()
	if kv.maxraftstate > 0 && kv.persister.RaftStateSize() > kv.maxraftstate {
		snapshot := kv.getSnapshot()
		Debug(dSnap, "G%d/S%d creates snapshot index=%d raftBytes=%d snapshotBytes=%d config=%d keys=%d", kv.gid, kv.me, msg.CommandIndex, kv.persister.RaftStateSize(), len(snapshot), kv.config.Num, len(kv.data))
		go func(i int) {
			kv.rf.Snapshot(i, snapshot)
		}(msg.CommandIndex)
	}
	kv.mu.Unlock()
}
