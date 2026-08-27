# Raft test logs

Use the scripts in `scripts/` to execute the Raft tests:

- `scripts/run-raft-2a.sh`
- `scripts/run-raft-2b.sh`
- `scripts/run-raft-2c.sh`

Each run writes a timestamped log file into this directory. The 2B and 2C
scripts also accept an optional Go test regular expression, for example:

```bash
./scripts/run-raft-2b.sh '^TestBasicAgree2B$'
./scripts/run-raft-2c.sh '^TestPersist12C$'
```

Set `RAFT_LOG_TOPICS` to a comma-separated list such as
`LOG1,LOG2,CMIT,PERS` when running `go test` directly to select log topics.

Generated `*.log` files are intentionally ignored by Git.
