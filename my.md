# my

- go build -o ./bin/mrcoordinator ./main/mrcoordinator.go
- go build -o ./bin/mrworker ./main/mrworker.go
- go build -o ./bin/mrsequential ./main/mrsequential.go

- go build -o ./bin/wc ./mrapps/wc.go
- go build -o ./bin/indexer ./mrapps/indexer.go

- go run ./main/mrcoordinator.go
- go run ./mrapps/wc.go


ibqo@ibqodeMBP src % go test ./raft
ok      6.5840/raft     302.763s

运行并编译测试（这是课程默认用法）：
- raft：go test ./raft
- kvraft：go test ./kvraft
- shardctrler：go test ./shardctrler

python3 -m pip install --user typer rich

raft % 
VERBOSE=1 dtest -n 2000 -p 5 -s  2A 这是raft运行2000次命令行，我怎么运行？

raft % 这是raft目录下 
VERBOSE=1 dtest -n 1000 -p 100 -s  2B
VERBOSE=1 dtest -n 10000 -p 100 -s  2C

kvraft目录
VERBOSE=1 dtest -n 500 -p 20 -s 3A
VERBOSE=1 dtest -n 500 -p 20 -s 3B
shardctrler目录
VERBOSE=1 dtest -n 500 -p 50 -s 4A

ibqo@ibqodeMBP mit-6.824_fravenx % cd src/raft 
ibqo@ibqodeMBP raft % export VERBOSE=1; python3 ./dtest -n 2 -p 2 -s 2A

备用方案（不用 dtest）
- 直接用 go test 连续运行（无并发控制、日志汇总）：
  - VERBOSE=1 go test -run 2A -count=2000
- 开启竞态检测：
  - VERBOSE=1 go test -run 2A -count=2000 -race

https://github.com/wdidada126/6.5840-golabs-2023

Lab of MIT 6.824 2023 所有lab 稳定通过一万次以上 All labs stably passed 10,000 times

'go version'
go version go1.16.6 linux/amd64
低版本go不行

cd src/main
./test-mr.sh

```shell
ibqodeMBP:main ibqo$ ./test-mr.sh
go version go1.17.13 darwin/amd64
*** Turning off -race since it may not work on a Mac
with  go version go1.17.13 darwin/amd64
*** Cannot find timeout command; proceeding without timeouts.
# command-line-arguments
ld: warning: -no_pie is deprecated when targeting new OS versions
# command-line-arguments
ld: warning: -no_pie is deprecated when targeting new OS versions
*** Starting wc test.
2025/11/27 22:52:18 rpc.Register: method "Done" has 1 input parameters; needs exactly three
--- wc test: PASS
*** Starting indexer test.
2025/11/27 22:52:25 rpc.Register: method "Done" has 1 input parameters; needs exactly three
--- indexer test: PASS
*** Starting map parallelism test.
2025/11/27 22:52:28 rpc.Register: method "Done" has 1 input parameters; needs exactly three
--- map parallelism test: PASS
*** Starting reduce parallelism test.
2025/11/27 22:52:36 rpc.Register: method "Done" has 1 input parameters; needs exactly three
--- reduce parallelism test: PASS
*** Starting job count test.
2025/11/27 22:52:45 rpc.Register: method "Done" has 1 input parameters; needs exactly three
--- job count test: PASS
*** Starting early exit test.
2025/11/27 22:53:02 rpc.Register: method "Done" has 1 input parameters; needs exactly three
--- early exit test: PASS
*** Starting crash test.
2025/11/27 22:53:11 rpc.Register: method "Done" has 1 input parameters; needs exactly three
--- crash test: PASS
*** PASSED ALL TESTS
```
