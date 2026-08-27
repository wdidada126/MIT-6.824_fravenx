# my

## mac
go 1.17不行，换go版本
export GOROOT="/Users/ibqo/development/go"
export PATH="/Users/ibqo/development/go/bin:$PATH"

```
ibqodeMBP:mit-6.824_fravenx ibqo$  export PATH="/Users/ibqo/development/go/bin:$PATH"
ibqodeMBP:mit-6.824_fravenx ibqo$ go version
go version go1.24.6 darwin/amd64
ibqodeMBP:mit-6.824_fravenx ibqo$ go env
AR='ar'
CC='clang'
CGO_CFLAGS='-O2 -g'
CGO_CPPFLAGS=''
CGO_CXXFLAGS='-O2 -g'
CGO_ENABLED='1'
CGO_FFLAGS='-O2 -g'
CGO_LDFLAGS='-O2 -g'
CXX='clang++'
GCCGO='gccgo'
GO111MODULE='on'
GOAMD64='v1'
GOARCH='amd64'
GOAUTH='netrc'
GOBIN=''
GOCACHE='/Users/ibqo/Library/Caches/go-build'
GOCACHEPROG=''
GODEBUG=''
GOENV='/Users/ibqo/Library/Application Support/go/env'
GOEXE=''
GOEXPERIMENT=''
GOFIPS140='off'
GOFLAGS=''
GOGCCFLAGS='-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -ffile-prefix-map=/var/folders/vk/_xdm46qn66182z1cy9k7qvr80000gn/T/go-build3066368003=/tmp/go-build -gno-record-gcc-switches -fno-common'
GOHOSTARCH='amd64'
GOHOSTOS='darwin'
GOINSECURE=''
GOMOD='/dev/null'
GOMODCACHE='/Users/ibqo/go/pkg/mod'
GONOPROXY=''
GONOSUMDB=''
GOOS='darwin'
GOPATH='/Users/ibqo/go'
GOPRIVATE=''
GOPROXY='https://goproxy.cn,direct'
GOROOT='/usr/local/go'
GOSUMDB=''
GOTELEMETRY='local'
GOTELEMETRYDIR='/Users/ibqo/Library/Application Support/go/telemetry'
GOTMPDIR=''
GOTOOLCHAIN=''
GOTOOLDIR='/usr/local/go/pkg/tool/darwin_amd64'
GOVCS=''
GOVERSION='go1.24.6'
GOWORK=''
PKG_CONFIG='pkg-config'
ibqodeMBP:mit-6.824_fravenx ibqo$
```


## doc script
```
- go build -o ./bin/mrcoordinator ./main/mrcoordinator.go
- go build -o ./bin/mrworker ./main/mrworker.go
- go build -o ./bin/mrsequential ./main/mrsequential.go

- go build -o ./bin/wc ./mrapps/wc.go
- go build -o ./bin/indexer ./mrapps/indexer.go

- go run ./main/mrcoordinator.go
- go run ./mrapps/wc.go
```

```
ibqo@ibqodeMBP src % go test ./raft
ok      6.5840/raft     302.763s
```

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


```shell
'go version'
go version go1.16.6 linux/amd64
```

低版本go不行

```shell
cd src/main
./test-mr.sh
```


```shell
ibqo@ibqodeMBP src % go test ./kvraft
ok      6.5840/kvraft   393.121s
```


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


```shell
cd src/labrpc
go test -v
=== RUN   TestTypes
--- PASS: TestTypes (0.00s)
=== RUN   TestDisconnect
--- PASS: TestDisconnect (0.00s)
=== RUN   TestCounts
--- PASS: TestCounts (0.00s)
=== RUN   TestBytes
--- PASS: TestBytes (0.00s)
=== RUN   TestConcurrentMany
--- PASS: TestConcurrentMany (0.00s)
=== RUN   TestUnreliable
--- PASS: TestUnreliable (0.03s)
=== RUN   TestConcurrentOne
--- PASS: TestConcurrentOne (0.00s)
=== RUN   TestRegression1
--- PASS: TestRegression1 (0.10s)
=== RUN   TestKilled
--- PASS: TestKilled (1.10s)
=== RUN   TestBenchmark
1.26843181s for 100000
--- PASS: TestBenchmark (1.27s)
PASS
ok      6.5840/labrpc   2.511s
```


```shell
cd src/labgob
go test -v
```

```shell
cd src/shardctrler
wdidada@LAPTOP-wdidada:/mnt/d/develops/git/github/go/MIT-6.824_fravenx/src/shardctrler$ go test -v
=== RUN   TestBasic4A
Test: Basic leave/join ...
doJoin %v at S%d 0xc0002da9f0 1
doJoin %v at S%d 0xc0002daff0 2
doJoin %v at S%d 0xc000314f90 0
doJoin %v at S%d 0xc00038f5f0 1
doJoin %v at S%d 0xc000421140 0
doJoin %v at S%d 0xc00038fbf0 2
doLeave %v at S%d [1/1]0xc000444608 1
doLeave %v at S%d [1/1]0xc000444768 0
doLeave %v at S%d [1/1]0xc00020b648 2
doLeave %v at S%d [1/1]0xc00020a248 1
doLeave %v at S%d [1/1]0xc00020a3a8 2
doLeave %v at S%d [1/1]0xc00011a668 0
  ... Passed
Test: Historical queries ...
doJoin %v at S%d 0xc0001170e0 0
doJoin %v at S%d 0xc000117140 0
doLeave %v at S%d [1/1]0xc00020b0c8 0
doLeave %v at S%d [1/1]0xc00020b0d0 0
doJoin %v at S%d 0xc00033a5d0 1
doJoin %v at S%d 0xc00033a630 1
doLeave %v at S%d [1/1]0xc00020bd38 1
doLeave %v at S%d [1/1]0xc00020bd40 1
  ... Passed
Test: Move ...
doJoin %v at S%d 0xc000486e40 2
doJoin %v at S%d 0xc000487590 0
doJoin %v at S%d 0xc000486ea0 2
doLeave %v at S%d [1/1]0xc00011b108 2
doLeave %v at S%d [1/1]0xc00011b110 2
doJoin %v at S%d 0xc0003a97d0 0
doJoin %v at S%d 0xc000487b90 2
doJoin %v at S%d 0xc0003a9dd0 2
doJoin %v at S%d 0xc0003a8c90 1
doJoin %v at S%d 0xc000465470 1
doMove shard = %d,gid = %d at S%d
 0 503 0
doMove shard = %d,gid = %d at S%d
 0 503 1
doMove shard = %d,gid = %d at S%d
 0 503 2
doMove shard = %d,gid = %d at S%d
 1 503 0
doMove shard = %d,gid = %d at S%d
 1 503 2
doMove shard = %d,gid = %d at S%d
 1 503 1
doMove shard = %d,gid = %d at S%d
 2 503 0
doMove shard = %d,gid = %d at S%d
 2 503 1
doMove shard = %d,gid = %d at S%d
 2 503 2
doMove shard = %d,gid = %d at S%d
 3 503 0
doMove shard = %d,gid = %d at S%d
 3 503 2
doMove shard = %d,gid = %d at S%d
 3 503 1
doMove shard = %d,gid = %d at S%d
 4 503 2
doMove shard = %d,gid = %d at S%d
 4 503 0
doMove shard = %d,gid = %d at S%d
 4 503 1
doMove shard = %d,gid = %d at S%d
 5 504 0
doMove shard = %d,gid = %d at S%d
 5 504 1
doMove shard = %d,gid = %d at S%d
 5 504 2
doMove shard = %d,gid = %d at S%d
 6 504 0
doMove shard = %d,gid = %d at S%d
 6 504 2
doMove shard = %d,gid = %d at S%d
 6 504 1
doMove shard = %d,gid = %d at S%d
 7 504 0
doMove shard = %d,gid = %d at S%d
 7 504 2
doMove shard = %d,gid = %d at S%d
 7 504 1
doMove shard = %d,gid = %d at S%d
 8 504 0
doMove shard = %d,gid = %d at S%d
 8 504 2
doMove shard = %d,gid = %d at S%d
 8 504 1
doMove shard = %d,gid = %d at S%d
 9 504 0
doMove shard = %d,gid = %d at S%d
 9 504 2
doMove shard = %d,gid = %d at S%d
 9 504 1
doLeave %v at S%d [1/1]0xc000445178 0
doLeave %v at S%d [1/1]0xc0004452e8 2
doLeave %v at S%d [1/1]0xc00020b7e8 1
doLeave %v at S%d [1/1]0xc00020b978 0
  ... Passed
doLeave %v at S%d [1Test: Concurrent leave/join ...
/1]0xc00020bae8 2
doLeave %v at S%d [1/1]0xc0004455e8 1
doJoin %v at S%d 0xc000501830 2
doJoin %v at S%d 0xc0003a9b90 2
doJoin %v at S%d 0xc000162570 1
doJoin %v at S%d 0xc0001625a0 1
doJoin %v at S%d 0xc00043f050 2
doJoin %v at S%d 0xc0004653e0 2
doJoin %v at S%d 0xc000465410 2
doJoin %v at S%d 0xc000501170 0
doJoin %v at S%d 0xc0003a9530 0
doJoin %v at S%d 0xc0001625d0 1
doJoin %v at S%d 0xc000162600 1
doJoin %v at S%d 0xc000162630 1
doJoin %v at S%d 0xc000162660 1
doJoin %v at S%d 0xc0002e8810 1
doJoin %v at S%d 0xc0002e8840 1
doJoin %v at S%d 0xc00043e9c0 0
doJoin %v at S%d 0xc000421290 0
doJoin %v at S%d 0xc00036a1b0 2
doJoin %v at S%d 0xc00036aff0 2
doJoin %v at S%d 0xc0003aea50 2
doJoin %v at S%d 0xc0001f2900 2
doJoin %v at S%d 0xc000464db0 0
doJoin %v at S%d 0xc000421740 0
doJoin %v at S%d 0xc0001160c0 1
doJoin %v at S%d 0xc00036a960 0
doJoin %v at S%d 0xc000558d80 0
doJoin %v at S%d 0xc000487b30 0
doJoin %v at S%d 0xc000315a70 2
doJoin %v at S%d 0xc00038e090 2
doJoin %v at S%d 0xc00038f230 2
doJoin %v at S%d 0xc00038f260 2
doJoin %v at S%d 0xc00038f290 2
doJoin %v at S%d 0xc0001d53b0 2
doJoin %v at S%d 0xc0001d53e0 2
doJoin %v at S%d 0xc0001d4210 0
doJoin %v at S%d 0xc00033a120 1
doJoin %v at S%d 0xc00033a150 1
doJoin %v at S%d 0xc0003151a0 0
doJoin %v at S%d 0xc00033ac60 1
doJoin %v at S%d 0xc00033ac90 1
doJoin %v at S%d 0xc00038f7d0 1
doJoin %v at S%d 0xc00038fd10 1
doJoin %v at S%d 0xc00038fd40 1
doJoin %v at S%d 0xc0002dafc0 0
doJoin %v at S%d 0xc0000dbf20 0
doJoin %v at S%d 0xc0002da0c0 0
doJoin %v at S%d 0xc000255620 0
doJoin %v at S%d 0xc0001d4600 0
doLeave %v at S%d [1/1]0xc000444e28 0
doLeave %v at S%d [1/1]0xc000193668 1
doJoin %v at S%d 0xc000558540 0
doJoin %v at S%d 0xc000559ec0 1
doLeave %v at S%d [1/1]0xc000193680 1
doJoin %v at S%d 0xc000559ef0 1
doLeave %v at S%d [1/1]0xc000445018 2
doJoin %v at S%d 0xc000558c90 2
doLeave %v at S%d [1/1]0xc000445548 2
doJoin %v at S%d 0xc0000dbd10 2
doJoin %v at S%d 0xc0001f27b0 1
doJoin %v at S%d 0xc00036a8d0 2
doJoin %v at S%d 0xc0002e8660 2
doLeave %v at S%d [1/1]0xc000444340 2
doLeave %v at S%d [1/1]0xc00040c428 2
doLeave %v at S%d [1/1]0xc00040cc98 0
doJoin %v at S%d 0xc000464bd0 0
doJoin %v at S%d 0xc0003a9cb0 0
doJoin %v at S%d 0xc0003afe00 0
doLeave %v at S%d [1/1]0xc00040c9b8 2
doLeave %v at S%d [1/1]0xc00040c9c0 2
doLeave %v at S%d [1/1]0xc00040d028 0
doLeave %v at S%d [1/1]0xc00040c168 0
doJoin %v at S%d 0xc0001f3410 1
doLeave %v at S%d [1/1]0xc000010588 1
doLeave %v at S%d [1/1]0xc000444528 0
doLeave %v at S%d [1/1]0xc00040d258 0
doLeave %v at S%d [1/1]0xc000444c88 2
doLeave %v at S%d [1/1]0xc000010590 1
doLeave %v at S%d [1/1]0xc000192038 0
doLeave %v at S%d [1/1]0xc000192798 1
doLeave %v at S%d [1/1]0xc0004448b0 1
doLeave %v at S%d [1/1]0xc000011228 0
doLeave %v at S%d [1/1]0xc00040ce48 1
doLeave %v at S%d [1/1]0xc00040da28 1
doLeave %v at S%d [1/1]0xc00040d778 2
doLeave %v at S%d [1/1]0xc000444038 0
doLeave %v at S%d [1/1]0xc0001922c8 0
doLeave %v at S%d [1/1]0xc000444258 2
doLeave %v at S%d [1/1]0xc00040c2a0 2
doLeave %v at S%d [1/1]0xc0000104a8 1
doLeave %v at S%d [1/1]0xc00040c628 1
  ... Passed
Test: Minimal transfers after joins ...
doJoin %v at S%d 0xc0001f3e60 0
doJoin %v at S%d 0xc00033a5a0 2
doJoin %v at S%d 0xc0002db110 1
doJoin %v at S%d 0xc0002db980 0
doJoin %v at S%d 0xc00038e0f0 2
doJoin %v at S%d 0xc000420420 1
doJoin %v at S%d 0xc00038f500 0
doJoin %v at S%d 0xc000420d20 2
doJoin %v at S%d 0xc0001d5500 1
doJoin %v at S%d 0xc0001d5ce0 0
doJoin %v at S%d 0xc0001d45a0 1
doJoin %v at S%d 0xc00033a750 2
doJoin %v at S%d 0xc00038ea20 0
doJoin %v at S%d 0xc00038f080 2
doJoin %v at S%d 0xc00033b320 1
  ... Passed
Test: Minimal transfers after leaves ...
doLeave %v at S%d [1/1]0xc0004446e8 0
doLeave %v at S%d [1/1]0xc000192cf8 2
doLeave %v at S%d [1/1]0xc000444988 1
doLeave %v at S%d [1/1]0xc00040c778 0
doLeave %v at S%d [1/1]0xc00040c9d8 2
doLeave %v at S%d [1/1]0xc000444df8 1
doLeave %v at S%d [1/1]0xc000193138 0
doLeave %v at S%d [1/1]0xc000193388 2
doLeave %v at S%d [1/1]0xc00040ce28 1
doLeave %v at S%d [1/1]0xc0004451e8 0
doLeave %v at S%d [1/1]0xc000010fc8 1
doLeave %v at S%d [1/1]0xc000445488 2
doLeave %v at S%d [1/1]0xc0001939b8 0
doLeave %v at S%d [1/1]0xc000193bf8 2
doLeave %v at S%d [1/1]0xc00040d200 1
  ... Passed
--- PASS: TestBasic4A (1.64s)
=== RUN   TestMulti4A
Test: Multi-group join/leave ...
doJoin %v at S%d 0xc0005010e0 0
doJoin %v at S%d 0xc00036b740 1
doJoin %v at S%d 0xc000465740 2
doJoin %v at S%d 0xc0003c09f0 0
doJoin %v at S%d 0xc0003c0ff0 2
doJoin %v at S%d 0xc000314a50 1
doLeave %v at S%d [2/2]0xc00020a690 0
doLeave %v at S%d [2/2]0xc00020a800 2
doLeave %v at S%d [2/2]0xc00020aaf0 1
doLeave %v at S%d [1/1]0xc000011458 0
doLeave %v at S%d   ... Passed
[1Test: Concurrent multi leave/join ...
/1]0xc0000115c8 2
doLeave %v at S%d [1/1]0xc000193118 1
doJoin %v at S%d 0xc000500a80 2
doJoin %v at S%d 0xc0003ae150 2
doJoin %v at S%d 0xc000330060 2
doJoin %v at S%d 0xc000331020 1
doJoin %v at S%d 0xc000331080 1
doJoin %v at S%d 0xc0003310e0 1
doJoin %v at S%d 0xc000500450 0
doJoin %v at S%d 0xc0000ff980 0
doJoin %v at S%d 0xc000465710 0
doJoin %v at S%d 0xc0003af7d0 0
doJoin %v at S%d 0xc0001f36b0 0
doJoin %v at S%d 0xc000117b60 0
doJoin %v at S%d 0xc000331bf0 0
doJoin %v at S%d 0xc000331140 1
doJoin %v at S%d 0xc0003311a0 1
doJoin %v at S%d 0xc0003308a0 1
doJoin %v at S%d 0xc000330900 1
doJoin %v at S%d 0xc000178780 2
doJoin %v at S%d 0xc0001f3d40 2
doJoin %v at S%d 0xc000540a50 2
doJoin %v at S%d 0xc0003afad0 0
doLeave %v at S%d [2/2]0xc0004444a0 0
doJoin %v at S%d 0xc0003141b0 2
doJoin %v at S%d 0xc000254570 1
doLeave %v at S%d [2/2]0xc000444e40 1
doJoin %v at S%d 0xc0002545d0 1
doLeave %v at S%d [2/2]0xc0004455b0 1
doLeave %v at S%d [2/2]0xc0004455c0 1
doLeave %v at S%d [2/2]0xc0004455d0 1
doJoin %v at S%d 0xc000421830 1
doJoin %v at S%d 0xc0000dbd40 0
doLeave %v at S%d [2/2]0xc000192b20 0
doLeave %v at S%d [2/2]0xc000594dd0 0
doJoin %v at S%d 0xc0001796e0 2
doLeave %v at S%d [2/2]0xc000444ac0 2
doJoin %v at S%d 0xc0002e8270 2
doLeave %v at S%d [2/2]0xc000193a60 1
doLeave %v at S%d [2/2]0xc00040c0f0 1
doLeave %v at S%d [2/2]0xc000192e40 2
doLeave %v at S%d [2/2]0xc000445280 0
doJoin %v at S%d 0xc000178120 0
doLeave %v at S%d [2/2]0xc0005950c0 2
doLeave %v at S%d [2/2]0xc00011bc00 0
doLeave %v at S%d [2/2]0xc000445910 2
doLeave %v at S%d [2/2]0xc000595a20 0
doJoin %v at S%d 0xc0001d5140 2
doLeave %v at S%d [2/2]0xc00011be60 2
doLeave %v at S%d [2/2]0xc000595c90 2
doLeave %v at S%d [2/2]0xc000010780 0
doLeave %v at S%d [2/2]0xc00040c040 0
doLeave %v at S%d [2/2]0xc00020ab90 1
doLeave %v at S%d [2/2]0xc00020aba0 1
doLeave %v at S%d [2/2]0xc00020abc0 1
doLeave %v at S%d [2/2]0xc00040c220 0
doLeave %v at S%d [2/2]0xc0000105e0 2
doLeave %v at S%d [2/2]0xc0000105f0 2
doLeave %v at S%d [2/2]0xc000011320 1
doLeave %v at S%d [2/2]0xc000010ca0 2
doLeave %v at S%d [2/2]0xc00060c500 0
doLeave %v at S%d [2/2]0xc00020b130 2
  ... Passed
Test: Minimal transfers after multijoins ...
doJoin %v at S%d 0xc0001d5cb0 0
doJoin %v at S%d 0xc0004205d0 2
doJoin %v at S%d 0xc0005581e0 1
  ... Passed
Test: Minimal transfers after multileaves ...
doLeave %v at S%d [5/5]0xc0005d4240 0
doLeave %v at S%d [5/5]0xc000457e00 2
doLeave %v at S%d [5/5]0xc0005d4600 1
  ... Passed
Test: Check Same config on servers ...
  ... Passed
--- PASS: TestMulti4A (1.29s)
PASS
ok      6.5840/shardctrler      2.929s
```

```
cd src/shardkv
go test -v
```

```
wdidada@LAPTOP-wdidada:/mnt/d/develops/git/github/go/MIT-6.824_fravenx/src/shardctrler$ cd src/shardkv
go test -v

bash: cd: src/shardkv:

=== RUN   TestBasic4A
Test: Basic leave/join ...
doJoin %v at S%d 0xc00030f290 1
doJoin %v at S%d 0xc00030f890 2
doJoin %v at S%d 0xc0001eb4d0 0
doJoin %v at S%d 0xc000329950 1
doJoin %v at S%d 0xc000329f50 2
doJoin %v at S%d 0xc00039d560 0
doLeave %v at S%d [1/1]0xc000011588 1
doLeave %v at S%d [1/1]0xc000376018 2
doLeave %v at S%d [1/1]0xc0000116e8 0
doLeave %v at S%d [1/1]0xc000285898 1
doLeave %v at S%d [1/1]0xc0000101c8 0
doLeave %v at S%d [1/1]0xc000376138 2
  ... Passed
Test: Historical queries ...
doJoin %v at S%d 0xc00023fd40 0
doJoin %v at S%d 0xc00023fda0 0
doLeave %v at S%d [1/1]0xc0002847b8 0
doLeave %v at S%d [1/1]0xc0002847c0 0
doJoin %v at S%d 0xc00009a7e0 1
doJoin %v at S%d 0xc00009a840 1
doLeave %v at S%d [1/1]0xc000285198 1
doLeave %v at S%d [1/1]0xc0002851a0 1
  ... Passed
Test: Move ...
doJoin %v at S%d 0xc0003431d0 0
doJoin %v at S%d 0xc0003429c0 2
doJoin %v at S%d 0xc000342a20 2
doLeave %v at S%d [1/1]0xc000377378 2
doLeave %v at S%d [1/1]0xc000377380 2
doJoin %v at S%d 0xc0001eb860 2
doJoin %v at S%d 0xc0003437d0 1
doJoin %v at S%d 0xc0004a4330 0
doJoin %v at S%d 0xc0004a4ba0 2
doJoin %v at S%d 0xc00039ca50 1
doMove shard = %d,gid = %d at S%d
 0 503 0
doMove shard = %d,gid = %d at S%d
 0 503 2
doMove shard = %d,gid = %d at S%d
 0 503 1
doMove shard = %d,gid = %d at S%d
 1 503 0
doMove shard = %d,gid = %d at S%d
 1 503 2
doMove shard = %d,gid = %d at S%d
 1 503 1
doMove shard = %d,gid = %d at S%d
 2 503 0
doMove shard = %d,gid = %d at S%d
 2 503 2
doMove shard = %d,gid = %d at S%d
 2 503 1
doMove shard = %d,gid = %d at S%d
 3 503 0
doMove shard = %d,gid = %d at S%d
 3 503 2
doMove shard = %d,gid = %d at S%d
 3 503 1
doMove shard = %d,gid = %d at S%d
 4 503 0
doMove shard = %d,gid = %d at S%d
 4 503 2
doMove shard = %d,gid = %d at S%d
 4 503 1
doMove shard = %d,gid = %d at S%d
 5 504 0
doMove shard = %d,gid = %d at S%d
 5 504 1
doMove shard = %d,gid = %d at S%d
 5 504 2
doMove shard = %d,gid = %d at S%d
 6 504 0
doMove shard = %d,gid = %d at S%d
 6 504 2
doMove shard = %d,gid = %d at S%d
 6 504 1
doMove shard = %d,gid = %d at S%d
 7 504 0
doMove shard = %d,gid = %d at S%d
 7 504 2
doMove shard = %d,gid = %d at S%d
 7 504 1
doMove shard = %d,gid = %d at S%d
 8 504 0
doMove shard = %d,gid = %d at S%d
 8 504 2
doMove shard = %d,gid = %d at S%d
 8 504 1
doMove shard = %d,gid = %d at S%d
 9 504 0
doMove shard = %d,gid = %d at S%d
 9 504 2
doMove shard = %d,gid = %d at S%d
 9 504 1
doLeave %v at S%d [1/1]0xc000119658 0
doLeave %v at S%d [1/1]0xc0001197c8 2
doLeave %v at S%d [1/1]0xc00027cdd8 1
doLeave %v at S%d [1/1]0xc00043d518 0
  ... Passed
Test: Concurrent leave/join ...
doLeave %v at S%d [1/1]0xc00043d688 1
doLeave %v at S%d [1/1]0xc000011428 2
doJoin %v at S%d 0xc0002f2450 1
doJoin %v at S%d 0xc0002f2480 1
doJoin %v at S%d 0xc0002f24b0 1
doJoin %v at S%d 0xc0002f24e0 1
doJoin %v at S%d 0xc0002f2510 1
doJoin %v at S%d 0xc0002f2540 1
doJoin %v at S%d 0xc0000cbc80 0
doJoin %v at S%d 0xc000343ad0 0
doJoin %v at S%d 0xc0002f2570 1
doJoin %v at S%d 0xc00048a330 2
doJoin %v at S%d 0xc0003b61b0 2
doJoin %v at S%d 0xc00043e7e0 0
doJoin %v at S%d 0xc0003294d0 0
doJoin %v at S%d 0xc00043f170 1
doJoin %v at S%d 0xc0004a4ab0 2
doJoin %v at S%d 0xc00043fbf0 1
doJoin %v at S%d 0xc0004a4ae0 2
doJoin %v at S%d 0xc0003b79b0 2
doJoin %v at S%d 0xc00048ab70 2
doJoin %v at S%d 0xc00043e690 2
doJoin %v at S%d 0xc0000fd290 2
doJoin %v at S%d 0xc00027f530 1
doJoin %v at S%d 0xc0003b7080 0
doJoin %v at S%d 0xc0004a4540 0
doJoin %v at S%d 0xc00027f560 1
doJoin %v at S%d 0xc0002f3740 2
doJoin %v at S%d 0xc0004a8d20 2
doJoin %v at S%d 0xc000329860 0
doJoin %v at S%d 0xc0003436b0 0
doJoin %v at S%d 0xc00027efc0 2
doJoin %v at S%d 0xc0001d4270 2
doJoin %v at S%d 0xc00048ba40 0
doJoin %v at S%d 0xc00009b830 1
doJoin %v at S%d 0xc0000fc960 0
doJoin %v at S%d 0xc00027e900 0
doJoin %v at S%d 0xc00009ac90 0
doJoin %v at S%d 0xc00030ed50 1
doJoin %v at S%d 0xc00030ed80 1
doJoin %v at S%d 0xc00030edb0 1
doJoin %v at S%d 0xc00030ede0 1
doJoin %v at S%d 0xc00030ee10 1
doLeave %v at S%d [1/1]0xc000010818 1
doLeave %v at S%d [1/1]0xc000376b50 1
doJoin %v at S%d 0xc00026e5d0 0
doJoin %v at S%d 0xc0003d2ae0 2
doJoin %v at S%d 0xc0003d2c90 2
doJoin %v at S%d 0xc0000caae0 2
doJoin %v at S%d 0xc0000cab10 2
doJoin %v at S%d 0xc0001eb5f0 2
doLeave %v at S%d [1/1]0xc00027c6f8 2
doJoin %v at S%d 0xc00043eae0 1
doJoin %v at S%d 0xc0002f2b10 1
doLeave %v at S%d [1/1]0xc0001187f0 2
doJoin %v at S%d 0xc00030e1b0 0
doJoin %v at S%d 0xc00048b620 2
doJoin %v at S%d 0xc0004a59b0 2
doJoin %v at S%d 0xc0001ea210 0
doJoin %v at S%d 0xc0001d57d0 0
doJoin %v at S%d 0xc00023f830 0
doLeave %v at S%d [1/1]0xc000010508 0
doJoin %v at S%d 0xc000383260 2
doJoin %v at S%d 0xc000382900 1
doLeave %v at S%d [1/1]0xc0001185d8 0
doJoin %v at S%d 0xc0001eaf00 0
doLeave %v at S%d [1/1]0xc000010498 2
doJoin %v at S%d 0xc0003d3dd0 0
doLeave %v at S%d [1/1]0xc000376768 2
doJoin %v at S%d 0xc0002f3a40 0
doLeave %v at S%d [1/1]0xc000377578 0
doLeave %v at S%d [1/1]0xc00027c800 2
doLeave %v at S%d [1/1]0xc000118ef0 2
doLeave %v at S%d [1/1]0xc000118ef8 2
doLeave %v at S%d [1/1]0xc000376558 0
doLeave %v at S%d [1/1]0xc0001186b8 0
doLeave %v at S%d [1/1]0xc000377418 2
doLeave %v at S%d [1/1]0xc000376078 0
doLeave %v at S%d [1/1]0xc000376158 0
doLeave %v at S%d [1/1]0xc0003771d8 0
doLeave %v at S%d [1/1]0xc000376298 1
doLeave %v at S%d [1/1]0xc000118c48 1
doLeave %v at S%d [1/1]0xc000118c50 1
doLeave %v at S%d [1/1]0xc000118c58 1
doLeave %v at S%d [1/1]0xc000376ec0 1
doLeave %v at S%d [1/1]0xc000011278 1
doLeave %v at S%d [1/1]0xc000011458 0
doLeave %v at S%d [1/1]0xc000119d68 2
doLeave %v at S%d [1/1]0xc0003779b8 0
doLeave %v at S%d [1/1]0xc000010200 2
doLeave %v at S%d [1/1]0xc0001182f8 1
doLeave %v at S%d [1/1]0xc000118300 1
  ... Passed
Test: Minimal transfers after joins ...
doJoin %v at S%d 0xc00027e840 0
doJoin %v at S%d 0xc00027ef00 2
doJoin %v at S%d 0xc0003d2150 1
doJoin %v at S%d 0xc000382f90 0
doJoin %v at S%d 0xc0003835c0 2
doJoin %v at S%d 0xc0003d3440 1
doJoin %v at S%d 0xc000129260 0
doJoin %v at S%d 0xc00030e750 2
doJoin %v at S%d 0xc00027fd10 1
doJoin %v at S%d 0xc0000caa80 0
doJoin %v at S%d 0xc00030fb00 2
doJoin %v at S%d 0xc0000cb0e0 1
doJoin %v at S%d 0xc00016fbf0 0
doJoin %v at S%d 0xc0001eb5c0 1
doJoin %v at S%d 0xc00048a180 2
  ... Passed
Test: Minimal transfers after leaves ...
doLeave %v at S%d [1/1]0xc00027c848 0
doLeave %v at S%d [1/1]0xc00027ca88 2
doLeave %v at S%d [1/1]0xc00043c348 1
doLeave %v at S%d [1/1]0xc000010528 0
doLeave %v at S%d [1/1]0xc00043c7d8 1
doLeave %v at S%d [1/1]0xc000376858 2
doLeave %v at S%d [1/1]0xc0000107c8 0
doLeave %v at S%d [1/1]0xc00043ca98 1
doLeave %v at S%d [1/1]0xc000010a20 2
doLeave %v at S%d [1/1]0xc000010cb8 0
doLeave %v at S%d [1/1]0xc000376be8 1
doLeave %v at S%d [1/1]0xc000010f20 2
doLeave %v at S%d [1/1]0xc00043d038 0
doLeave %v at S%d [1/1]0xc00043d280 2
doLeave %v at S%d [1/1]0xc000011258 1
  ... Passed
--- PASS: TestBasic4A (1.50s)
=== RUN   TestMulti4A
Test: Multi-group join/leave ...
doJoin %v at S%d 0xc00043e1b0 1
doJoin %v at S%d 0xc00043e8d0 2
doJoin %v at S%d 0xc00032e870 0
doJoin %v at S%d 0xc00026f1d0 1
doJoin %v at S%d 0xc000328600 0
doJoin %v at S%d 0xc00026f7d0 2
doLeave %v at S%d [2/2]0xc000376230 1
doLeave %v at S%d [2/2]0xc0003763b0 2
doLeave %v at S%d [2/2]0xc00043c5e0 0
doLeave %v at S%d [1/1]0xc00027cd38 1
doLeave %v at S%d [1/1]0xc00027ceb8 2
  ... Passed
Test: Concurrent multi leave/join ...
doLeave %v at S%d [1/1]0xc000376dd8 0
doJoin %v at S%d 0xc00030f0e0 1
doJoin %v at S%d 0xc00030f800 2
doJoin %v at S%d 0xc00032e330 2
doJoin %v at S%d 0xc00032fbc0 2
doJoin %v at S%d 0xc00032fc20 2
doJoin %v at S%d 0xc0000fc780 2
doJoin %v at S%d 0xc00032ef30 0
doLeave %v at S%d [2/2]0xc0001183f0 2
doJoin %v at S%d 0xc0000ca360 1
doJoin %v at S%d 0xc00043e330 1
doJoin %v at S%d 0xc00023e7e0 1
doJoin %v at S%d 0xc00048b770 1
doLeave %v at S%d [2/2]0xc00043ddc0 1
doJoin %v at S%d 0xc00039cc90 1
doJoin %v at S%d 0xc00032eff0 0
doJoin %v at S%d 0xc000492240 1
doJoin %v at S%d 0xc000328bd0 2
doJoin %v at S%d 0xc00032f050 0
doJoin %v at S%d 0xc00023f590 0
doJoin %v at S%d 0xc0004380f0 0
doLeave %v at S%d [2/2]0xc000118910 0
doJoin %v at S%d 0xc000438150 0
doJoin %v at S%d 0xc0002e4f00 2
doLeave %v at S%d [2/2]0xc000118280 1
doJoin %v at S%d 0xc000439200 1
doLeave %v at S%d [2/2]0xc0003c89b0 1
doJoin %v at S%d 0xc0002e4360 1
doLeave %v at S%d [2/2]0xc0001191e0 2
doLeave %v at S%d [2/2]0xc000118f30 1
doJoin %v at S%d 0xc0004381b0 0
doJoin %v at S%d 0xc0004a4900 2
doLeave %v at S%d [2/2]0xc000284d90 0
doLeave %v at S%d [2/2]0xc000119210 2
doJoin %v at S%d 0xc0001a5740 0
doJoin %v at S%d 0xc0004a4a20 2
doLeave %v at S%d [2/2]0xc000284dc0 0
doJoin %v at S%d 0xc0001a57a0 0
doLeave %v at S%d [2/2]0xc0003c9730 0
doLeave %v at S%d [2/2]0xc000285e30 0
doLeave %v at S%d [2/2]0xc000119250 2
doJoin %v at S%d 0xc00039c300 0
doLeave %v at S%d [2/2]0xc000285600 2
doLeave %v at S%d [2/2]0xc0002853b0 1
doJoin %v at S%d 0xc0003d36e0 2
doJoin %v at S%d 0xc0001ea7e0 1
doLeave %v at S%d [2/2]0xc00043c2d0 2
doLeave %v at S%d [2/2]0xc00043c040 1
doLeave %v at S%d [2/2]0xc0000101f0 0
doLeave %v at S%d [2/2]0xc000010370 1
doLeave %v at S%d [2/2]0xc000376890 1
doLeave %v at S%d [2/2]0xc00027c340 1
doLeave %v at S%d [2/2]0xc00027c750 0
doLeave %v at S%d [2/2]0xc00027c760 0
doLeave %v at S%d [2/2]0xc000010680 2
doLeave %v at S%d [2/2]0xc00027c5d0 2
doLeave %v at S%d [2/2]0xc00027c5e0 2
doLeave %v at S%d [2/2]0xc00027c770 0
doLeave %v at S%d [2/2]0xc0000114f0 1
doLeave %v at S%d [2/2]0xc000377a90 0
doLeave %v at S%d [2/2]0xc000011770 2
  ... Passed
Test: Minimal transfers after multijoins ...
doJoin %v at S%d 0xc00030f200 1
doJoin %v at S%d 0xc00030f9b0 2
doJoin %v at S%d 0xc0004a9260 0
  ... Passed
Test: Minimal transfers after multileaves ...
doLeave %v at S%d [5/5]0xc000178c90 1
doLeave %v at S%d [5/5]0xc00021b650 2
doLeave %v at S%d [5/5]0xc000282900 0
  ... Passed
Test: Check Same config on servers ...
  ... Passed
--- PASS: TestMulti4A (1.42s)
PASS
ok      6.5840/shardctrler      2.922s
```

```
cd src/kvraft
3A 3B
wdidada@LAPTOP-wdidada:/mnt/d/develops/git/github/go/MIT-6.824_fravenx/src/kvraft$ go test -v
--- PASS: TestSnapshotRecoverManyClients3B (21.18s)
=== RUN   TestSnapshotUnreliable3B
Test: unreliable net, snapshots, many clients (3B) ...
  ... Passed --  16.9  5 12072 1183
--- PASS: TestSnapshotUnreliable3B (16.91s)
=== RUN   TestSnapshotUnreliableRecover3B
Test: unreliable net, restarts, snapshots, many clients (3B) ...
  ... Passed --  20.8  5 12697 1202
--- PASS: TestSnapshotUnreliableRecover3B (20.83s)
=== RUN   TestSnapshotUnreliableRecoverConcurrentPartition3B
Test: unreliable net, restarts, partitions, snapshots, many clients (3B) ...
  ... Passed --  28.1  5  9762  822
--- PASS: TestSnapshotUnreliableRecoverConcurrentPartition3B (28.12s)
=== RUN   TestSnapshotUnreliableRecoverConcurrentPartitionLinearizable3B
Test: unreliable net, restarts, partitions, snapshots, random keys, many clients (3B) ...
  ... Passed --  31.9  7 28562 1904
--- PASS: TestSnapshotUnreliableRecoverConcurrentPartitionLinearizable3B (31.89s)
PASS
ok      6.5840/kvraft   391.454s
```

---

# 测试报告（2026-08-23，Go 1.24.2 linux/amd64）

环境：Ubuntu 24.04 (WSL2)；Go 安装于 ~/.local/go（go1.24.2）；go.mod 为 go 1.15，Go 1.24 下编译通过。
测试命令：`cd src && go test ./raft ./kvraft ./shardctrler ./shardkv`；Lab1 用 `cd src/main && ./test-mr.sh`。

## 测试结果（全部 PASS）

| Lab | 测试 | 结果 | 耗时 |
|---|---|---|---|
| Lab1 | test-mr.sh（wc/indexer/map并行/reduce并行/jobcount/early exit/crash） | PASS 7/7 | ~8min |
| Lab2 | raft 2A 选举 | PASS | 16.6s |
| Lab2 | raft 2B 日志复制 | PASS | 35.8s |
| Lab2 | raft 2C 持久化 | PASS | 118.6s |
| Lab2 | raft 2D 快照 | PASS | 139.1s |
| Lab3 | kvraft 3A | PASS | 246.9s |
| Lab3 | kvraft 3B | PASS | 142.2s |
| Lab4a | shardctrler 4A | PASS | 3.0s |
| Lab4b | shardkv 基本功能x4 | PASS | 17.7s |
| Lab4b | shardkv 并发x3 | PASS | 32.9s |
| Lab4b | shardkv 不可靠x3 | PASS | 19.4s |
| Lab4b | shardkv Challengex3 | PASS | 25.8s |

说明：
- raft 全量（2A~2D）一起跑 5 分钟跑不完（2C+2D 单测约 4.3 分钟），拆开逐个跑全部通过；CI 工作流中已给足超时。
- go vet 仅有课程自带测试代码的惯用写法警告，非错误。
- CI 已配置：.github/workflows/test.yml（push/PR 触发，Go 1.24.x，分 5 步跑 lab1~lab4b）。

---

# 各 Lab 知识点与源码位置

## Lab1 MapReduce（mr）

知识点：
1. MapReduce 模型：Map 按 key 哈希分 nReduce 个中间桶，Reduce 按 key 聚合；中间文件先写临时文件再原子重命名，避免读到半成品。
2. Coordinator 调度：任务三态（idle/in-progress/completed），worker 通过 RPC 拉取任务（pull 模式）；超时未完成任务被重置，实现 crash 容错。
3. worker 主循环：循环拉任务 -> 执行 -> 汇报结果。

源码位置：
- src/mr/coordinator.go：MakeCoordinator L168（初始化）、GetTask L44（任务分发）、ReportTask L91（任务汇报）、resetExpired L125（超时重置）、allMapsDone L139 / allReducesDone L148、Done L159
- src/mr/worker.go：Worker L39（主循环）、runMap L94（map 任务）、runReduce L160（reduce 任务）、ihash L32（key 分桶）
- src/mr/rpc.go：RPC 消息与 Task 结构定义

## Lab2 Raft（raft）

知识点：
1. 领导选举：term + 随机超时（150-300ms）+ 过半投票 + 先到先得；票数过半转 leader。
2. 日志复制：AppendEntries 做 prevLogIndex/prevLogTerm 一致性检查；leader 用 nextIndex/matchIndex 推进；firstIndex 快速回退优化。
3. 安全性：只 commit 当前 term 的日志；投票时比较日志新旧（term 大者/日志长者为新）。
4. 持久化（2C）：currentTerm/votedFor/log 三个字段 persist。
5. 快照（2D）：Snapshot 截断 log；落后 follower 通过 InstallSnapshot 直接发快照。
6. apply 单例：单独 applier 协程按序发 applyCh（sync.Cond 通知），保证不重不漏、无并发乱序。

源码位置（src/raft/raft.go）：
- 选举：ticker L590（超时触发选举）、RequestVote L245（投票处理）
- 日志复制：AppendEntries L272、handleAppendEntriesReply L761（回退推进）、updateCommitIndex L554（commit 推进）
- 持久化：persist L113 / readPersist L126
- 快照：Snapshot L158、InstallSnapshot L346、installSnapshot L804
- 其他：Start L430（提交入口）、applier L842、sendHeartbeat L680 / sendHeartbeatTo L724、Make L897

## Lab3 KVRaft（kvraft）

知识点：
1. 线性一致性：所有读写都进 raft 日志串行执行，从 applyCh 取出后执行并回给等待的 RPC。
2. 幂等去重：clientId + seqNo 去重表，写操作重放不重复执行（isRepeated）。
3. leader 切换：client 遇 ErrWrongLeader / RPC 失败换 leader 重试。
4. 快照：raft 状态超过 maxraftstate 阈值时持久化 server 状态（data + 去重表），重启读快照恢复。

源码位置：
- src/kvraft/server.go：Get L39、PutAppend L92（写日志）、isRepeated L165（幂等去重）、execute L173（apply 后执行）、getSnapshot L251 / readSnapshot L260、StartKVServer L290
- src/kvraft/client.go：Get L43、PutAppend L75（带 clientId+seqNo 重试）

## Lab4a ShardCtrler（shardctrler）

知识点：
1. 配置版本：configs 数组按 Num 递增保存每届配置快照，Query 支持历史版本查询。
2. Rebalance：维护 gid->shard 集合，取分片最多/最少的组，相差 >1 时成批搬移，使各组分片尽量均衡（minimal transfer）。
3. 状态机复制：Join/Leave/Move/Query 全部走 raft 日志执行。

源码位置：
- src/shardctrler/server.go：Join L38、Leave L73、Move L109、Query L144、doJoin L218（join+rebalance）、doLeave L251（leave+rebalance）、doMove L292、isRepeated L304
- src/shardctrler/executer.go：execute L5（apply 循环）、applyMsg L15、reply L57

## Lab4b ShardKV（shardkv）

知识点：
1. 分片状态机：Working/Missing/Adding，非 Working 的分片拒绝读写请求（checkShard）。
2. 配置逐版应用：全部 shard 就绪才拉取下一版配置（readyForNewConfig），避免跨版本直接迁移导致数据不一致。
3. 分片迁移：PushShardRPC 携带 data + 去重表 + 源/目标 gid + 配置 Num，重复接收不破坏数据。
4. 垃圾回收：DeleteShardRPC 通知源集群删除已迁走的 shard（GC），否则分片残留会导致 2B 测试失败。
5. 快照持久化：config + lastConfig + shard 状态 + data + 去重表。

源码位置（src/shardkv/server.go）：
- checkShard L57（shard 状态检查）、Get L68 / PutAppend L127、PushShard L181（接收迁移）、DeleteShard L241（GC）
- updateConfig L323（轮询拉配置）、updateShardsState L344（状态标记）、readyForNewConfig L314
- preparePushShardArgs L368（构造迁移数据）、sendPushShardRPC L401 / sendDeleteShardRPC L413 / sendShards L433
- getSnapshot L450 / readSnapshot L463
- src/shardkv/client.go：key2shard L20（key->shard 映射）、Get L64 / PutAppend L98（先查配置定位 group）

---

# 项目 README（博客结构版）

> MIT 6.824/6.5840 分布式系统课程全 4 个 Lab 完整实现（MapReduce / Raft / KVRaft / ShardKV，含 4B Challenges），已通过全部测试并接入 GitHub Actions CI。

## 一、项目概览

| Lab | 模块 | 目录 | 核心内容 |
|---|---|---|---|
| Lab1 | MapReduce | src/mr | Coordinator 任务调度 + Worker 执行，容 crash |
| Lab2 | Raft 共识 | src/raft | 选举、日志复制、持久化、快照 |
| Lab3 | KVRaft | src/kvraft | 线性一致 KV，幂等去重，快照恢复 |
| Lab4a | ShardCtrler | src/shardctrler | 分片配置管理 + 自动 Rebalance |
| Lab4b | ShardKV | src/shardkv | 多组集群分片迁移 + GC + Challenges |

## 二、各 Lab 核心设计（要点）

- **Lab1 MapReduce**：Map 输出按 key 哈希分 nReduce 个中间文件；Coordinator 维护任务三态（idle/in-progress/completed），worker 通过 RPC 拉取任务（pull 模式）；任务超时（10s）自动重置以容 worker crash；中间文件先写临时文件再原子重命名。
- **Lab2 Raft**：随机超时（150-300ms）触发选举，term 单调递增，得票过半转 leader；AppendEntries 做 prevLogIndex/prevLogTerm 一致性检查，冲突段用 FirstIndex 快速回退；只 commit 当前 term 的日志；currentTerm/votedFor/log 持久化；快照压缩日志（2D）。
- **Lab3 KVRaft**：所有读写都进 raft 日志串行执行，从 applyCh 取命令后应用并回给等待的 RPC；clientId+seqNo 去重表保证幂等；raft 状态超阈值时持久化 server 快照（data + 去重表）。
- **Lab4a ShardCtrler**：configs 按 Num 递增保存配置快照；Join/Leave/Move/Query 全走 raft 日志；Rebalance 策略：max-min 分片差 >1 时每次搬一半，保证各组负载均衡且 minimal transfer。
- **Lab4b ShardKV**：每 shard 三态（Working/Missing/Adding），非 Working 拒绝请求；配置逐版应用（全部就绪才拉下一版）；PushShardRPC 迁移数据+去重表，DeleteShardRPC 通知源组 GC。

## 三、架构与核心时序

### 3.1 Raft 选举与日志复制

```mermaid
sequenceDiagram
    autonumber
    participant F1 as Follower A
    participant C as Candidate/Leader
    participant F2 as Follower B
    Note over F1: electionTimeout 超时
    F1->>F1: term++, 投自己一票
    F1->>C: RequestVote(term, lastLogIndex, lastLogTerm)
    F1->>F2: RequestVote(term, lastLogIndex, lastLogTerm)
    C-->>F1: 同意（日志不比自己旧）
    F2-->>F1: 同意
    Note over F1: 得票过半 → 转 Leader
    F1->>C: AppendEntries(心跳, term, prevLogIndex, entries)
    F1->>F2: AppendEntries(心跳, term, prevLogIndex, entries)
    C-->>F1: Success（prevLogIndex/prevLogTerm 匹配）
    F2-->>F1: Success
    Note over F1: commitIndex 推进 → applier 协程下发 applyCh
```

### 3.2 KVRaft：Client 请求 → Apply 回调完整流程

```mermaid
sequenceDiagram
    autonumber
    participant C as Clerk(Client)
    participant S as KVServer(Leader)
    participant R as Raft(Leader)
    C->>S: PutAppend/Get RPC (clientId, seqNo)
    S->>R: rf.Start(op)
    R-->>S: index, term, isLeader=true
    S->>S: 注册 waitCh[index]
    R->>R: 日志复制到多数派并 commit
    R-->>S: applyCh <- ApplyMsg{Command:op}
    S->>S: execute(): 幂等检查(isRepeated) → 更新 data/table
    S-->>C: 通过 waitCh[index] 返回结果
    Note over C: 若 ErrWrongLeader/超时 → 换一个服务器重试
```

### 3.3 ShardKV：配置更新 + 分片迁移握手（Push + GC）

```mermaid
sequenceDiagram
    autonumber
    participant M as ShardCtrler(Master)
    participant A as Group A(源, gid=1)
    participant B as Group B(目标, gid=2)
    A->>M: Query(configNum+1) 轮询
    M-->>A: 新配置 Num=2（部分分片归 gid=2）
    A->>A: Start(config) 写入日志并 apply
    A->>A: updateShardsState: 失去的分片标记 MISSING
    A->>B: PushShard(Shards, Data, Table, Num, Src, Dst)
    Note over B: 校验 config.Num 关系 + shard 状态
    B->>B: Start(PushShard) 写入日志，apply 后置 WORKING
    B-->>A: OK
    A->>B: DeleteShard(Keys, Num, Dst)
    Note over B: 源分片数据 GC，raft 状态不膨胀
    B-->>A: OK
```

## 四、测试结果与性能数据

环境：Ubuntu 24.04 (WSL2) / Go 1.24.2 / 单机。命令：`cd src && go test ./raft ./kvraft ./shardctrler ./shardkv`。

### 4.1 单轮耗时（全部 PASS）

| 测试 | 耗时 | 测试 | 耗时 |
|---|---|---|---|
| raft 2A | 16.6s | kvraft 3A | 246.9s |
| raft 2B | 35.8s | kvraft 3B | 142.2s |
| raft 2C | 118.6s | shardctrler 4A | 3.0s |
| raft 2D | 139.1s | shardkv 4B 全量 | ~96s |
| lab1 test-mr.sh | ~8min | shardkv Challenge | ~26s |

### 4.2 稳定性（防 flaky 验证，连续 3 次）

| 测试 | 3 连跑总耗时 | 通过率 |
|---|---|---|
| raft 2C | 358.9s | 3/3 PASS |
| raft 2D | 411.7s | 3/3 PASS |
| shardkv Challenge（3 项） | 83.8s | 3/3 PASS |

结论：2C/2D 持久化与快照路径稳定，无 flaky；Challenge 在 -count=3 下全部通过。

### 4.3 4B Challenge 测试场景说明

- **TestChallenge1Delete（分片删除/GC）**：`make_config(t, 3, false, 1)` —— maxraftstate=1 强制每条日志后做快照。写 30 个 key（各 1KB）后反复 `join(1)→leave(0)→join(2)→leave(1)→join(0)→leave(2)` 两轮。校验各 group 的 RaftStateSize：若迁走的分片数据未 GC，raft 状态会持续膨胀超过限制 → 失败。**核心验证点：DeleteShard 必须真正删除数据。**
- **TestChallenge2Unaffected（未受影响分片访问）**：网络不可靠 + 快照。JOIN 100 → 写数据 → JOIN 101（部分 shard 迁到 101）→ **KILL 100** → LEAVE 100。此时 101 永远无法从 100 迁到数据，但校验 101 自有 shard 仍可正常读写。**核心验证点：部分迁移失败不得影响已拥有分片。**
- **TestChallenge2Partial（部分迁移）**：三组同时 JOIN 后写数据，在迁移未全部完成时校验未受影响分片可访问。**核心验证点：分片状态机（Working/Adding）正确处理未完成迁移。**

## 五、踩坑记录（重点）

### 5.1 Lab2 Raft

1. **firstIndex（快照截断）与 log 物理下标的关系**：`compactTo`（raft.go L499）截断后 `log[0]` 变为哨兵项 `Entry{Index: lastIncludedIndex, Term: lastIncludedTerm}`，后续 `rf.log[0].Index` 不再是 1。**所有按逻辑 index 取日志的操作都必须走 `getByIndex`（L523，二分查找返回物理下标）**，直接 `rf.log[index - firstIndex]` 在快照后必然越界或取错。典型错误出现在 `handleAppendEntriesReply` L774：回退前必须用 `rf.getByIndex(rf.nextIndex[i]-1)` 校验 leader 自身日志，防止用过期回复覆盖 nextIndex。
2. **AppendEntries 的 prevLogIndex 命中快照截断点**（L291）：follower 侧用 `getByIndex(args.PrevLogIndex)` 判断：
   - `PrevLogIndex == log[0].Index`（恰好是哨兵项）→ exists=true，只需校验 Term 匹配（哨兵项 Term 即 lastIncludedTerm）；
   - `PrevLogIndex < log[0].Index`（follower 落后于快照点）→ exists=false，回复 `Success=false, FirstIndex=-1`，leader 侧用 term 扫描快速回退（L779-784）；
   - 若直接写 `rf.log[PrevLogIndex - log[0].Index]`，负数下标会 panic，且语义错误。
3. **append 日志的覆盖语义**（L314-325）：逐条比较 `cmpCommand(log[i+k], args.Entries[i-1])`，不匹配则截断 `rf.log[:i+k]` 再追加，否则跳过（已存在）；`i+k < len(rf.log)` 判断要基于物理下标。
4. **apply 必须单协程 + 条件变量**：`applier`（L842）持有锁循环，`commitIndex > lastApplied` 时组装 ApplyMsg 批量发送，发前释放锁（防 channel 阻塞卡死整个 Raft）；commit 推进后用 `applyCond.Signal()`（L335）唤醒。若在 RPC 处理协程里直接发 applyCh，一旦上层消费慢会阻塞心跳/选举。
5. **commit 推进的安全性**：`updateCommitIndex`（L554）从 `commitIndex+1` 起找**当前 term** 的日志过半才推进（防上一 term 的日志被"间接提交"后又因新 leader 覆盖导致回滚）；`getByIndex` 不存在（已截断）则停止。

### 5.2 Lab3 KVRaft

1. **去重表在快照恢复后正确重建**：快照只 encode `table`（clientId→最大 seqNo）+ `data`（kvraft/server.go `getSnapshot` L251）。`readSnapshot`（L260）解码后整体替换两个 map，**顺序不能反**（先 table 后 data，与 encode 顺序一致）。坑：若快照里 table 落后于最近已应用的 op，重启后该 op 会被重放 —— 设计上依赖"快照总是在状态变更完成后、`bytes` 超阈值时立刻生成"（execute L227-235），且幂等表保证重放无害。
2. **重复写操作也要处理快照**（L198-211）：`isRepeated` 命中的 op 虽然不执行，但**仍可能触发快照**（否则 `bytes` 一直涨），且必须 `continue` 跳过后续，不能把 `kv.data` 再写一遍。
3. **Get 也更新 table**（L183-185）：Get 无副作用但同样记录 seqNo，用于后续重试幂等；注意 Get 返回时若 `kv.table[clientId] == op.SeqNo` 才回值（L187），防止读到旧值。
4. **waitCh 通知必须 non-blocking**：`select { case kv.waitCh[clientId] <- op: default: }`（L190、L221），因为 RPC handler 可能已因 50ms 超时返回，通道满会卡死 apply 循环。

### 5.3 Lab4b ShardKV（含 Challenge）

1. **PushShard 幂等性——防止"重复迁移把新数据覆盖成旧数据"**（server.go L181-239），按 `config.Num` 与 `args.Num` 三路判断：
   - `config.Num < args.Num`（目标组还没到该配置）：**直接返回 OK 并丢弃，不写入日志**（L187-191）。否则日志里出现"未来配置的迁移"，apply 时会基于旧 config 覆盖数据；目标组靠后续 updateConfig 追赶 + 源组重发来完成真正迁移。
   - `config.Num > args.Num`（旧配置的迟到迁移）：直接 OK + 触发 DeleteShard 帮源组 GC（L193-198）。
   - `config.Num == args.Num` 且 shard 已 WORKING（重复迁移）：直接 OK + DeleteShard（L199-211），**绝不重写 data** —— 这是防覆盖的关键；只有处于 MISSING 才 `rf.Start(*args)` 写日志真正执行。
2. **DeleteShard 只带 Keys 而非整个 Data**（L416-421）：GC 时按 key 删除，防止误删迁移后目标组的新写入（challenge1 直接检查 raft 状态大小，若 GC 不彻底会膨胀超限）。
3. **updateConfig 轮询节奏**（L323-342）：只有 `readyForNewConfig()`（无 MISSING 分片）且是 leader 时才 Query(Num+1)；非 leader/未就绪则 sleep 跳过。若配置跳版（一次 Query 拿到最新版），旧迁移一律按"Num 不匹配"丢弃，不会污染。
4. **shard 状态机初始化**（L344-357）：`lastConfig.Num == 0`（首个配置）直接 return，避免把"上版归属别人"误标为 ADDING；状态迁移只在 `lastConfig.Num != 0` 后按 lastConfig→config 逐 shard 比较。
5. **迁移参数必须携带 Num/Src/Dst/Table**：`preparePushShardArgs`（L368）打包 data + `copyTable(kv.table)`（去重表随迁，否则目标组重启后会重放源组已执行的写） + Src/Dst server 地址 + Num。src/dst 从 `lastConfig.Groups` / `config.Groups` 取（L393/395），不能写死。

### 5.4 Lab1 MapReduce（简）

- 中间文件命名 `mr-X-Y`（X map 任务号，Y reduce 桶号），reduce 端必须按 key 排序后聚合；`ihash` 决定桶归属，map 端和 worker 端哈希函数必须一致。
- Coordinator 对"卡死任务"用 10s 超时重置（`resetExpired` coordinator.go L125），但**不能重置已完成任务**，否则重复分配导致重复输出；ReportTask 里按任务类型分别更新计数。
- `Done` 判定要等所有 reduce 完成（`allReducesDone` L148），且全部状态访问加锁。

### 5.5 其他小坑

- `shardctrler/server.go` L219/L253/L293 存在调试残留 `println("doJoin %v at S%d", ...)`——Go 内置 `println` 不支持格式化串，会原样打印字符串且无法取值，测试输出会乱；应改用 `fmt.Printf` 或删除（建议删除，不影响正确性）。
- `test-mr.sh` 原为 CRLF 行尾导致 Linux 下 shebang 失效，已通过 `.gitattributes`（`*.sh text eol=lf`）+ CI 中 sed 双保险修复。
- `go.mod` 声明 go 1.15，Go 1.24 下直接编译通过，无需改动。

1. 补充“踩坑记录”
每个 Lab 都有一些经典难题，您可以补充实际编码中遇到的坑和解决思路，这对读者更有参考价值。例如：

Lab2：firstIndex 与快照配合时 log 切片索引的边界计算（容易 out of range）

Lab2：AppendEntries 中 prevLogIndex 命中了快照截断点时的处理

Lab3：重复请求（duplicate request）去重表在快照恢复后如何正确重建

Lab4b：跨配置版本迁移时，如何避免“重复迁移导致数据被旧数据覆盖”（即 PushShard 的幂等性细节）

2. 增加架构图或时序图

Raft 选举/日志复制的核心时序

KVRaft 从 Client 发起请求到 Apply 回调的完整流程

ShardKV 配置更新 + 分片迁移的握手过程（Push + GC）

3. 整理性能数据
您已有耗时数据，可以补充：

Raft 2C/2D 持久化和快照的测试通过次数（是否稳定无 flaky）

4B Challenge 测试的具体场景（如网络分区 + 并发迁移同时发生）

4. 格式化为博客/项目 README
如果您打算把这份文档发布为技术博客或 GitHub 项目说明，我可以帮您：

重新组织成“项目概览 → 各 Lab 核心设计 → 测试结果 → 实现细节 → 遇到的坑”结构

---

# 如何运行 2A 测试并通过日志理解 Raft

这个项目不需要先启动服务。2A 测试会自动在进程内创建多个 Raft 节点、模拟网络断连并检查选举结果。

## 运行 2A 测试

进入 Go module 所在目录：

```bash
cd /Users/ibqo/Develop/git/github/golang/mit-6.824_fravenx/src
```

运行全部 2A：

```bash
go test ./raft -run '2A$' -count=1 -v
```

分别运行三个用例：

```bash
go test ./raft -run '^TestInitialElection2A$' -count=1 -v
go test ./raft -run '^TestReElection2A$' -count=1 -v
go test ./raft -run '^TestManyElections2A$' -count=1 -v
```

三个用例分别观察：

- `InitialElection`：3 个节点首次选出 Leader。
- `ReElection`：Leader 断网、重新选举、丢失多数派、恢复多数派。
- `ManyElections`：7 个节点反复随机断开和恢复。
- `-count=1`：禁止 Go 使用测试缓存，确保每次真的重新运行。

首个用例已经实际验证，当前代码可以通过。

## 打开 Raft 日志

日志由 `VERBOSE` 环境变量控制，定义在 `src/raft/util.go`。

```bash
VERBOSE=1 go test ./raft -run '^TestInitialElection2A$' -count=1 -v
```

保存并实时查看：

```bash
VERBOSE=1 go test ./raft -run '2A$' -count=1 -v 2>&1 \
  | tee /tmp/raft-2a.log
```

事后筛选关键事件：

```bash
rg 'VOTE|LEAD|TERM|TIMR|PASS|FAIL' /tmp/raft-2a.log
```

典型输出：

```text
005380 VOTE S1 become leader
```

含义是：

- `005380`：程序启动后约 538 ms，单位是 0.1 ms。
- `VOTE`：投票/选举主题。
- `S1`：编号为 1 的 Raft 节点。
- `become leader`：该节点获得多数票，成为 Leader。

测试末尾例如：

```text
... Passed --  3.5  3  62  16040  0
```

依次表示：耗时、节点数、RPC 数、RPC 字节数、达成一致的日志条数。2A 只测试选举，所以最后通常是 `0`。

## 当前日志为什么比较少

当前选举路径中的很多日志被注释了，例如：

- 收到、拒绝和同意投票：`src/raft/raft.go` 的 `RequestVote` 方法。
- 选举超时及发起选举：`src/raft/raft.go` 的 `ticker` 方法。
- 发送 `RequestVote`：`ticker` 中调用 `sendRequestVote` 的位置。
- 成为 Leader：统计到多数票之后设置 `rf.state = LEADER` 的位置。

为了学习，建议临时取消这些 `Debug(dVote, ...)` 行前面的注释。然后运行：

```bash
VERBOSE=1 go test ./raft -run '^TestReElection2A$' -count=1 -v
```

可以观察到下面的因果链：

```text
选举超时
→ currentTerm 增加
→ Follower 变成 Candidate
→ 给自己投票
→ 向其他节点发送 RequestVote
→ Follower 接受或拒绝
→ 获得多数票
→ Candidate 变成 Leader
→ Leader 发送心跳
→ 旧 Leader 重连后看到更大 Term，退回 Follower
```

不要一开始打开每一次心跳日志，因为心跳间隔只有 100 ms，会很吵。先关注 `Term`、`Candidate`、`Vote`、`Leader` 四类变化最容易理解。

## 并发与稳定性检查

用 race detector 检查锁和共享状态：

```bash
VERBOSE=1 go test -race ./raft -run '2A$' -count=1 -v
```

仓库还提供了批量压力测试器 `src/raft/dtest`：

```bash
cd raft
./dtest '2A$' -n 100 -p 4
```

保存所有轮次日志：

```bash
./dtest '2A$' -n 20 -p 4 -v -a -o /tmp/raft-2a-runs
```

学习时推荐先按 `InitialElection → ReElection → ManyElections` 的顺序单独运行，理解正确后再使用 `dtest` 压测。

## 2A日志解析

缩写	含义
TIMR	选举计时器超时（election timeout）
VOTE	投票相关事件
TERM	任期变更
LEAD	成为 leader
DROP	网络消息被故意丢弃（模拟网络分区/丢包）
节点编号：S0、S1、S2...（7个节点时到 S6）


Raft 论文中的术语	你的日志缩写
election timeout	TIMR
RequestVote RPC	VOTE
term	TERM
leader	LEAD
网络分区/丢包	DROP
论文只定义算法逻辑，不关心你怎么打印日志。

## Raft 概念与项目日志命名的区别

这些缩写和 `S0、S1...` 的显示格式，都是这个 Go 项目自己定义的日志约定，不是 Raft 算法强制规定的。

具体区分如下：

| 项目 | Raft 算法概念 | 当前代码的表示方式 |
|---|---|---|
| 选举超时 | 是 | 日志缩写为 `TIMR` |
| 投票 | 是 | 日志缩写为 `VOTE` |
| 任期 Term | 是 | 日志缩写为 `TERM` |
| Leader | 是 | 日志缩写为 `LEAD` |
| 网络丢包 | Raft 需要应对，但不规定日志 | 测试网络记录为 `DROP` |
| 节点唯一编号 | 算法需要区分节点 | 当前代码使用 `S0、S1、S2...` |

这些缩写定义在 `src/raft/util.go`：

```go
dDrop   logTopic = "DROP"
dLeader logTopic = "LEAD"
dTerm   logTopic = "TERM"
dTimer  logTopic = "TIMR"
dVote   logTopic = "VOTE"
```

`S0` 中的 `S` 也是日志字符串手动写出来的：

```go
Debug(dLeader, "S%d becomes leader ...", rf.me)
```

其中 `rf.me` 是当前节点在节点数组中的下标：

```go
me int
```

因此节点从 `0` 开始编号：

```text
rf.me = 0 → S0
rf.me = 1 → S1
rf.me = 2 → S2
```

节点数量由测试代码决定：

- `TestInitialElection2A`：`servers := 3`，所以是 `S0～S2`。
- `TestReElection2A`：`servers := 3`。
- `TestManyElections2A`：`servers := 7`，所以是 `S0～S6`。

这些定义可以在 `src/raft/test_test.go` 中看到。

总结来说：Raft 规定了节点、Term、投票、Candidate、Follower 和 Leader 等概念，但不规定日志缩写、节点名称，也不规定编号必须从 0 开始。换成 `Node-A`、`server-1` 或中文日志，都不会影响 Raft 算法。

---

# Raft 2B：日志复制、提交与应用

2A 解决“谁是 Leader”，2B 继续解决“Leader 怎样让多数节点保存相同的命令”。

## 运行和保存 2B 日志

在项目根目录运行全部 2B 测试：

```bash
./scripts/run-raft-2b.sh
```

日志会保存在：

```text
logs/raft/2b-时间戳.log
```

刚开始学习时，建议先单独运行基础一致性测试：

```bash
./scripts/run-raft-2b.sh '^TestBasicAgree2B$'
```

然后依次观察 Follower 掉线、重新加入和日志冲突回退：

```bash
./scripts/run-raft-2b.sh '^TestFollowerFailure2B$'
./scripts/run-raft-2b.sh '^TestRejoin2B$'
./scripts/run-raft-2b.sh '^TestBackup2B$'
```

脚本内部设置了：

```bash
VERBOSE=1
RAFT_LOG_TOPICS=TEST,LEAD,TERM,LOG1,LOG2,CMIT
```

如果直接使用 `go test`，可以这样保存日志：

```bash
cd src
VERBOSE=1 RAFT_LOG_TOPICS=LEAD,TERM,LOG1,LOG2,CMIT \
  go test ./raft -run '^TestBasicAgree2B$' -count=1 -v 2>&1 \
  | tee ../logs/raft/2b-basic.log
```

## 2B 新增日志主题

| 缩写 | 含义 |
|---|---|
| `LOG1` | Leader 接受 `Start()`，把命令加入自己的日志 |
| `LOG2` | AppendEntries、一致性检查、复制确认以及 `nextIndex` 回退 |
| `CMIT` | `commitIndex` 前进，或者日志通过 `applyCh` 应用 |
| `LEAD` | 节点成为 Leader |
| `TERM` | 任期或角色发生变化 |

空的心跳不会输出 `LOG2`，只有真正携带日志或者引起复制状态变化时才重点记录。

## 从日志观察一次完整复制

典型日志如下：

```text
LOG1 Leader S2 accepts Start(command=100): index=1 term=1
LOG2 S1 accepts entries [1..1] from leader S2; last index=1
LOG2 Leader S2 receives replication ack from S1: matchIndex=1 nextIndex=2
CMIT Leader S2 advances commitIndex 0 -> 1 in term 1
CMIT S2 applies index=1 command=100
CMIT S1 advances commitIndex 0 -> 1 from leader S2
CMIT S1 applies index=1 command=100
```

对应的算法过程是：

```text
测试器调用 Leader.Start(command)
→ Leader 将 Entry 追加到本地 log
→ Leader 通过 AppendEntries 复制给 Follower
→ Follower 检查 prevLogIndex 和 prevLogTerm
→ Follower 保存 Entry 并返回成功
→ Leader 更新该 Follower 的 matchIndex 和 nextIndex
→ 多数节点已经保存 Entry
→ Leader 推进 commitIndex
→ Leader 通过 applyCh 应用 Entry
→ 后续心跳携带 LeaderCommit
→ Follower 推进 commitIndex 并应用 Entry
```

注意：`Start()` 返回成功只代表当前节点认为自己是 Leader，并已经把命令加入本地日志，不代表命令已经提交。必须在多数节点确认后，才能看到 `CMIT ... advances commitIndex`。

## nextIndex 和 matchIndex

Leader 为每个 Follower 保存两个进度：

- `nextIndex[i]`：下一次希望发给节点 `i` 的日志索引。
- `matchIndex[i]`：已知节点 `i` 与 Leader 匹配的最大日志索引。

复制成功时可能看到：

```text
LOG2 Leader S0 receives replication ack from S2: matchIndex=8 nextIndex=9
```

日志冲突时可能看到：

```text
LOG2 S2 rejects AppendEntries from S0: prev index 8 has term 2, want 3
LOG2 Leader S0 backtracks nextIndex[2] 9 -> 5
```

含义是 Follower 的日志与 Leader 不一致，Leader 向前回退 `nextIndex`，从更早的位置重新复制。

## 2B 测试覆盖内容

| 测试 | 主要检查内容 |
|---|---|
| `TestBasicAgree2B` | 基本日志复制、提交和应用 |
| `TestRPCBytes2B` | RPC 传输字节数不能异常增长 |
| `TestFollowerFailure2B` | 一个 Follower 掉线后，多数派仍可提交 |
| `TestLeaderFailure2B` | Leader 掉线后重新选举并继续提交 |
| `TestFailAgree2B` | Follower 断开并重新加入后追赶日志 |
| `TestFailNoAgree2B` | 丢失多数派时不能提交 |
| `TestConcurrentStarts2B` | 并发调用 `Start()` 时日志索引保持正确 |
| `TestRejoin2B` | 旧 Leader 重新加入后的冲突日志修复 |
| `TestBackup2B` | Follower 日志差距很大时快速回退和追赶 |
| `TestCount2B` | 心跳及复制 RPC 数量不能过高 |

## 筛选 2B 日志

只看命令从写入到应用：

```bash
rg 'LOG1|CMIT' logs/raft/2b-*.log
```

只看冲突与回退：

```bash
rg 'rejects AppendEntries|backtracks nextIndex' logs/raft/2b-*.log
```

只看 Leader 和任期变化：

```bash
rg 'LEAD|TERM' logs/raft/2b-*.log
```

---

# Raft 2C：持久化、崩溃恢复与不可靠网络

2C 的目标是：节点崩溃并重启后，仍能恢复 Raft 的关键状态，不会在同一任期重复投票，也不会丢失已经保存的日志。

## 运行和保存 2C 日志

运行全部 2C 测试：

```bash
./scripts/run-raft-2c.sh
```

日志会保存在：

```text
logs/raft/2c-时间戳.log
```

完整 2C 包含 Figure 8、不可靠网络和 churn，日志很长。学习时应先运行最基础的持久化测试：

```bash
./scripts/run-raft-2c.sh '^TestPersist12C$'
```

然后逐步运行：

```bash
./scripts/run-raft-2c.sh '^TestPersist22C$'
./scripts/run-raft-2c.sh '^TestPersist32C$'
./scripts/run-raft-2c.sh '^TestFigure82C$'
./scripts/run-raft-2c.sh '^TestUnreliableAgree2C$'
```

脚本默认关注：

```bash
RAFT_LOG_TOPICS=TEST,LEAD,TERM,LOG1,CMIT,PERS
```

## 2C 新增日志主题

| 缩写 | 含义 |
|---|---|
| `PERS` | Raft 状态写入持久化存储，或者节点启动时恢复状态 |
| `LOG1` | Leader 接收新命令并追加本地日志 |
| `CMIT` | 提交或应用日志 |
| `TERM` | 节点发现新任期并更新状态 |
| `LEAD` | 新 Leader 当选，测试中也可能用它标记节点崩溃 |

## 哪些状态必须持久化

按照 Raft Figure 2，服务器必须持久化：

- `currentTerm`：当前已知的最大任期。
- `votedFor`：当前任期投给了哪个 Candidate。
- `log[]`：日志条目的 `Index`、`Term` 和 `Command`。

本项目为 2D 快照还会一起保存：

- `lastIncludedIndex`。
- `lastIncludedTerm`。

`commitIndex` 和 `lastApplied` 不属于 Raft Figure 2 要求的持久化状态。重启后可以根据 Leader 的 `LeaderCommit` 和重新复制的日志恢复提交进度。

## 从日志观察持久化与恢复

写入持久化状态：

```text
PERS S1 persists term=1 votedFor=1 log=[0..1] bytes=109
```

含义是：

- 节点：`S1`。
- 当前任期：`1`。
- 本任期投票给：`S1`，也就是自己。
- 当前持久化日志索引范围：`0..1`。
- 编码后的 Raft 状态大小：`109` 字节。

节点重启并恢复：

```text
PERS S1 restores term=1 votedFor=1 log=[0..1] bytes=109
```

这表示新的 Raft 实例从 `Persister` 中恢复了崩溃前的 Term、投票和日志，而不是从空状态开始。

一次典型的 2C 因果链：

```text
Leader 接收 command=11
→ 日志变更后调用 persist()
→ 多数派保存并提交 index=1
→ 节点崩溃，内存状态消失
→ 测试器保留 Persister 中的数据
→ Make() 创建新的 Raft 实例
→ readPersist() 恢复 term、votedFor 和 log
→ 节点重新选举或接收 Leader 心跳
→ 已恢复日志继续参与一致性检查和提交
```

重要原则是：需要持久化的状态发生变化后，必须先完成 `persist()`，再让其他协程或 RPC 观察到该状态。典型位置包括：

- `currentTerm` 增加或更新。
- `votedFor` 改变。
- Leader 在 `Start()` 中追加日志。
- Follower 在 `AppendEntries` 中增加、覆盖或截断日志。

## 2C 测试覆盖内容

| 测试 | 主要检查内容 |
|---|---|
| `TestPersist12C` | 基础崩溃、重启和日志恢复 |
| `TestPersist22C` | 多节点分区、重启后的持续一致性 |
| `TestPersist32C` | 分区 Leader/Follower 崩溃后的恢复 |
| `TestFigure82C` | Raft 论文 Figure 8 的复杂 Leader 变更场景 |
| `TestUnreliableAgree2C` | 丢包、延迟和重复 RPC 下的一致性 |
| `TestFigure8Unreliable2C` | Figure 8 与不可靠网络组合 |
| `TestReliableChurn2C` | 节点反复连接、断开、崩溃和重启 |
| `TestUnreliableChurn2C` | churn 再叠加不可靠网络 |

## 筛选 2C 日志

只看持久化与恢复：

```bash
rg ' PERS ' logs/raft/2c-*.log
```

只看恢复事件：

```bash
rg ' PERS .*restores' logs/raft/2c-*.log
```

同时查看崩溃、恢复和 Leader 变化：

```bash
rg 'crashes|restores| LEAD | TERM ' logs/raft/2c-*.log
```

查看某个节点，例如 `S1` 的完整轨迹：

```bash
rg ' S1 ' logs/raft/2c-*.log
```

完整 churn 日志适合验证稳定性，不适合第一次阅读。第一次学习优先看 `TestPersist12C` 的日志，再看 `TestPersist32C` 和 `TestFigure82C`。

---

# Lab 3A：基于 Raft 的容错 KV 服务

Lab 2 实现了 Raft 共识模块。Lab 3A 在 Raft 上增加一个键值状态机，对外提供 `Get`、`Put` 和 `Append`。

## 运行和保存 3A 日志

在项目根目录运行全部 3A 测试：

```bash
./scripts/run-lab3a.sh
```

单独运行某个测试：

```bash
./scripts/run-lab3a.sh '^TestBasic3A$'
./scripts/run-lab3a.sh '^TestOnePartition3A$'
./scripts/run-lab3a.sh '^TestUnreliableOneKey3A$'
```

日志保存到：

```text
logs/lab3/3a-时间戳.log
```

脚本默认日志主题：

```bash
RAFT_LOG_TOPICS=LEAD,TERM
KV_LOG_TOPICS=CLNT,DUPL
```

如果想观察每个副本如何 Apply 命令，可以临时打开 `KVOP`：

```bash
KV_LOG_TOPICS=CLNT,KVOP,DUPL \
  ./scripts/run-lab3a.sh '^TestBasic3A$'
```

`TestBasic3A` 自身会执行上千条命令。阅读时可先看前几十条：

```bash
sed -n '1,80p' logs/lab3/3a-*.log
```

## 3A 应用层日志主题

| 缩写 | 含义 |
|---|---|
| `CLNT` | KVServer 将客户端请求提交给 Raft，或者完成请求 |
| `KVOP` | KV 状态机从 `applyCh` 收到并执行命令 |
| `DUPL` | 根据 ClientId 和 SeqNo 检测并跳过重复请求 |
| `LEAD` | 底层 Raft 选出 Leader |
| `TERM` | 底层 Raft 任期或角色变化 |

## 一次 Put 的完整过程

典型日志：

```text
CLNT S3 proposes Put client=... seq=1 key="0" value="x" at index=1 term=1
KVOP S3 applies Put index=1 client=... seq=1 key="0" value="x"
CLNT S3 completes Put client=... seq=1 key="0"
KVOP S1 applies Put index=1 client=... seq=1 key="0" value="x"
KVOP S2 applies Put index=1 client=... seq=1 key="0" value="x"
```

对应流程：

```text
Clerk 生成 ClientId 和递增 SeqNo
→ Clerk 向它记住的 Leader 发送 PutAppend RPC
→ KVServer 把请求包装为 Op
→ KVServer 调用 rf.Start(Op)
→ Raft 将 Op 复制到多数节点并提交
→ 每个 KVServer 从 applyCh 收到相同 Op
→ 每个副本按相同顺序修改 data
→ 发起请求的 Leader 通过 waitCh 唤醒 RPC handler
→ RPC 返回 OK，Clerk 才增加 SeqNo
```

这里有两个不同层次：

- Raft 的 `log[]` 保存的是等待共识的 `Op`。
- KVServer 的 `data map[string]string` 是命令被提交后得到的状态机结果。

客户端看到 `OK` 时，命令已经经过 Raft 提交并由该 KVServer Apply，而不仅仅是写进 Leader 的本地日志。

## 为什么 Get 也要进入 Raft

`Get` 不修改键值，但当前实现仍将它包装成 `Op` 写入 Raft。这样可以让读取排在之前所有写操作之后，并确认处理它的节点仍是有效 Leader，从而支持线性一致读。

其路径是：

```text
Get RPC
→ rf.Start(Get Op)
→ 多数派提交
→ Apply Get
→ 从 data 中读取值并回复客户端
```

## 请求重试与去重

网络可能出现以下情况：

```text
请求已经提交
→ Leader 的回复丢失
→ Clerk 超时并向其他服务器重试同一个请求
```

如果再次执行 `Append`，值会被追加两次。因此 Clerk 为每个逻辑请求携带：

- `ClientId`：客户端唯一标识。
- `SeqNo`：该客户端单调递增的请求序号。

每个 KVServer 保存：

```go
table map[int64]int64 // ClientId -> 已执行的最大 SeqNo
```

当 `SeqNo <= table[ClientId]` 时，请求已经执行过，状态机不会再次修改数据，并输出 `DUPL` 日志。

这实现的是在客户端按顺序、一次只发出一个请求这一前提下的“重复请求只执行一次”。

## Leader 切换和 waitCh

RPC handler 调用 `Start()` 后不能立即返回成功，因为命令可能永远无法提交。它会创建一个带缓冲的 `waitCh`，等待 Apply 协程通知。

可能结果包括：

- Apply 到相同的 ClientId/SeqNo：返回 `OK`。
- 相同位置出现另一个 Leader 的命令：返回错误并让 Clerk 重试。
- 等待超时：返回 `TimeOut`，Clerk 更换服务器。
- `Start()` 发现不是 Leader：立即返回 `ErrWrongLeader`。

通道必须带缓冲，避免 RPC 已经超时退出后 Apply 协程永久阻塞。

## 3A 测试覆盖内容

| 测试类型 | 主要检查内容 |
|---|---|
| `TestBasic3A` | 单客户端 Put/Get/Append |
| `TestSpeed3A` | 请求完成速度 |
| `TestConcurrent3A` | 多客户端并发 |
| `TestUnreliable*3A` | 丢包、延迟和重复 RPC 下的去重与一致性 |
| `TestOnePartition3A` | 多数派可继续服务，少数派不能提交 |
| `TestManyPartitions*3A` | 网络分区反复变化 |
| `TestPersist*3A` | KVServer 与 Raft 反复重启 |
| `*Linearizable3A` | 使用 Porcupine 检查历史是否线性一致 |

## 筛选 3A 日志

查看 Leader 接收和完成请求：

```bash
rg ' CLNT ' logs/lab3/3a-*.log
```

只看重复请求：

```bash
rg ' DUPL ' logs/lab3/3a-*.log
```

跟踪某个客户端：

```bash
rg 'client=客户端编号' logs/lab3/3a-*.log
```

---

# Lab 3B：KV 状态机快照

Raft 日志会持续增长。3B 要求 KVServer 在 Raft 状态超过 `maxraftstate` 后生成快照，使 Raft 可以删除已经被状态机吸收的旧日志。

## 运行和保存 3B 日志

运行全部 3B：

```bash
./scripts/run-lab3b.sh
```

第一次学习最推荐运行 InstallSnapshot 用例：

```bash
./scripts/run-lab3b.sh '^TestSnapshotRPC3B$'
```

也可以观察崩溃恢复：

```bash
./scripts/run-lab3b.sh '^TestSnapshotRecover3B$'
```

日志保存到：

```text
logs/lab3/3b-时间戳.log
```

## 3B 日志主题

| 缩写 | 含义 |
|---|---|
| `SNAP` | KVServer 创建、恢复或安装快照；Raft 截断日志 |
| `DUPL` | 从快照恢复去重表后识别重复请求 |
| `LEAD` / `TERM` | 快照场景中的 Leader 和任期变化 |

典型日志：

```text
SNAP S1 creates snapshot at index=46 raftBytes=1052 snapshotBytes=309 keys=45 clients=2
SNAP S1 snapshots to 46 log[46 - 46]
SNAP S2 restores snapshot bytes=309 keys=45 clients=2
SNAP S2 installs snapshot index=46 term=2 bytes=309 keys=45 clients=2
```

含义：

1. KVServer S1 发现 Raft 状态超过阈值。
2. 它编码 KV 数据和客户端去重表，得到 309 字节的快照。
3. 它调用 `rf.Snapshot(46, snapshot)`，通知 Raft 索引 46 及之前的日志已不再需要。
4. 落后的 S2 无法通过普通 AppendEntries 获得已删除的日志。
5. Leader 改用 InstallSnapshot RPC 将快照发送给 S2。
6. S2 恢复 `data` 和 `table`，然后从快照之后的日志继续追赶。

## 快照必须包含什么

当前 KV 快照编码：

```text
data   ：所有键值
table  ：ClientId -> 最大 SeqNo
```

去重表必须一起保存。如果只恢复键值而不恢复去重表，客户端重试崩溃前已经完成的 Append 时会再次执行，破坏 exactly-once 语义。

Raft 自己还会持久化：

```text
currentTerm、votedFor、剩余 log、lastIncludedIndex、lastIncludedTerm
```

KV 快照和 Raft 状态需要原子保存，防止崩溃后两者对应不上。

## 快照创建和恢复路径

```text
KVServer Apply index=N
→ 检查 persister.RaftStateSize()
→ 编码 data 和 table
→ rf.Snapshot(N, bytes)
→ Raft 记录快照边界并截断 log
```

重启路径：

```text
StartKVServer
→ persister.ReadSnapshot()
→ readSnapshot()
→ 恢复 data 和 table
→ 启动 Apply 协程继续处理新日志
```

落后 Follower 路径：

```text
Leader 发现 nextIndex[follower] <= lastIncludedIndex
→ 发送 InstallSnapshot
→ Follower 安装快照并更新提交/应用边界
→ 后续继续使用 AppendEntries
```

## 筛选 3B 日志

只看创建快照：

```bash
rg 'creates snapshot' logs/lab3/3b-*.log
```

只看恢复或安装：

```bash
rg 'restores snapshot|installs snapshot' logs/lab3/3b-*.log
```

查看日志截断范围：

```bash
rg 'snapshots to' logs/lab3/3b-*.log
```

完整恢复测试会创建数千次快照，不适合第一次阅读。优先查看 `TestSnapshotRPC3B` 的精简日志。

---

# Lab 4A：Shard Controller 和配置重平衡

Lab 4 将键空间划分为 10 个 shard。ShardCtrler 负责维护“每个 shard 当前属于哪个复制组”的版本化配置。

## 运行和保存 4A 日志

运行全部 4A：

```bash
./scripts/run-lab4a.sh
```

单独运行基础配置测试：

```bash
./scripts/run-lab4a.sh '^TestBasic4A$'
```

日志保存到：

```text
logs/lab4/4a-时间戳.log
```

脚本默认关注：

```bash
RAFT_LOG_TOPICS=LEAD,TERM
CTR_LOG_TOPICS=CONF,CLNT,DUPL
```

## 4A 日志主题和节点名称

| 表示 | 含义 |
|---|---|
| `SC0`、`SC1`、`SC2` | 三个 ShardCtrler 副本 |
| `CLNT` | Controller Leader 将 Join/Leave/Move/Query 写入 Raft |
| `CONF` | Controller 状态机生成并应用新配置 |
| `DUPL` | 重复配置修改请求被去重 |

`SC` 是本项目日志使用的 Shard Controller 缩写，不是 Raft 论文规定的名称。

## Config 的结构

```go
type Config struct {
    Num    int
    Shards [10]int
    Groups map[int][]string
}
```

- `Num`：配置版本，从 0 单调增加。
- `Shards[i]`：shard i 当前所属的 GID。
- `Groups[gid]`：该复制组中的服务器列表。
- 配置 0 没有任何组，所有 shard 都属于无效 GID 0。

## Join、Leave、Move 和 Query

- `Join`：增加一个或多个复制组，并重新平衡 shard。
- `Leave`：删除复制组，并将它们的 shard 分配给剩余组。
- `Move`：明确把一个 shard 移动到指定 GID。
- `Query(-1)`：查询最新配置。
- `Query(num)`：查询历史配置。

修改操作必须通过 ShardCtrler 自己的 Raft 日志复制，使全部 Controller 副本按同一顺序产生相同配置。

## 从日志观察重平衡

```text
CONF SC2 applies Join: config=1 groups=[1] shards=[1 1 1 1 1 1 1 1 1 1]
CONF SC2 applies Join: config=2 groups=[1 2] shards=[2 2 2 2 2 1 1 1 1 1]
CONF SC2 applies Leave gids=[1]: config=3 groups=[2] shards=[2 2 2 2 2 2 2 2 2 2]
```

解释：

1. GID 1 首次加入，10 个 shard 全部给它。
2. GID 2 加入，两个组各获得 5 个 shard。
3. GID 1 离开，它原来的 shard 全部转给 GID 2。

每个 `CONF` 通常会在 `SC0/SC1/SC2` 各出现一次，因为同一个配置修改被三个 Controller 副本 Apply。

## 重平衡目标

当存在 N 个有效组时，算法需要满足：

- 任意两个组的 shard 数量之差不超过 1。
- 尽量少移动已经合理分配的 shard。
- 相同输入必须产生确定性的配置，不能依赖 Go map 的随机遍历顺序。

当前实现先对 GID 排序，再反复从 shard 最多的组移动到最少的组。排序很重要，否则不同 Controller 副本可能基于相同命令产生不同配置。

## 为什么保存所有历史配置

ShardKV 迁移 shard 时需要知道：

```text
上一个配置中 shard 属于谁
当前配置中 shard 属于谁
```

因此 Controller 保存 `configs[]`，以配置编号作为数组下标。`Query(num)` 必须能返回旧配置，不能只保留最新配置。

## 筛选 4A 日志

只看配置变化：

```bash
rg ' CONF ' logs/lab4/4a-*.log
```

查看某个 shard 的 Move：

```bash
rg 'applies Move shard=3' logs/lab4/4a-*.log
```

---

# Lab 4B：分片 KV、配置推进与 shard 迁移

4B 中存在多个 Raft 复制组。每个组只服务当前配置分配给自己的 shard；配置改变时，数据和去重信息必须安全地迁移到新组。

## 运行和保存 4B 日志

运行全部 4B 和 Challenge：

```bash
./scripts/run-lab4b.sh
```

第一次学习最推荐 Join/Leave：

```bash
./scripts/run-lab4b.sh '^TestJoinLeave$'
```

观察快照与迁移：

```bash
./scripts/run-lab4b.sh '^TestSnapshot$'
```

日志保存到：

```text
logs/lab4/4b-时间戳.log
```

脚本默认过滤掉多个 Raft 组的底层投票噪音，关注：

```bash
CTR_LOG_TOPICS=CONF
SHARDKV_LOG_TOPICS=CONF,MIGR,SNAP,DUPL
```

如需观察客户端操作：

```bash
SHARDKV_LOG_TOPICS=CONF,MIGR,SNAP,KVOP,DUPL \
  ./scripts/run-lab4b.sh '^TestJoinLeave$'
```

## 4B 节点名称与日志主题

`G100/S2` 表示：

- `G100`：GID 为 100 的复制组。
- `S2`：该组内编号为 2 的 Raft/KV 副本。

| 缩写 | 含义 |
|---|---|
| `CONF` | ShardCtrler 配置，或 ShardKV 提议/应用新配置 |
| `MIGR` | PushShard、安装数据、DeleteShard 和垃圾回收 |
| `KVOP` | 分片 KV 请求的路由、提议与 Apply |
| `SNAP` | 包含配置和 shard 状态的 ShardKV 快照 |
| `DUPL` | 客户端请求或迁移 RPC 的幂等处理 |

## key 如何映射到 shard

当前测试代码使用：

```text
shard = int(key[0]) % 10
```

客户端先计算 shard，再查看 `config.Shards[shard]` 找到 GID，然后尝试该组中的服务器。如果收到 `ErrWrongGroup`，客户端会向 ShardCtrler 查询最新配置并重新路由。

## 三种 shard 状态

每个 ShardKV 组维护：

| 状态 | 含义 |
|---|---|
| `WORKING` | shard 可以正常服务，或者该组无需等待它 |
| `MISSING` | 旧配置属于本组、新配置已转出；本组需要向新组发送数据 |
| `ADDING` | 新配置属于本组，但数据尚未从旧组到达 |

日志会直接列出状态：

```text
working=[5 6 7 8 9] missing=[0 1 2 3 4] adding=[]
working=[5 6 7 8 9] missing=[] adding=[0 1 2 3 4]
```

第一行通常属于源组，第二行属于目标组。

## 配置必须逐个推进

ShardKV 只查询 `currentConfig.Num + 1`，并且只有所有 shard 都回到 `WORKING` 时才接受下一个配置：

```text
配置 N 被 Raft 提交
→ 完成 N 引起的所有 shard 迁移和垃圾回收
→ 所有状态回到 WORKING
→ 查询并提交配置 N+1
```

不能从配置 1 直接跳到配置 3，否则无法确定配置 2 中每个 shard 的正确数据来源。

配置查询在单一后台循环中串行执行。不能周期性创建多个共享同一 Clerk 的 Query goroutine，否则它们会并发修改 Clerk 的 SeqNo，并破坏逐配置推进顺序。

## 一次完整 shard 迁移

下面是一条实际日志链：

```text
CONF G100/S2 applies config 1 -> 2 ... missing=[0 1 2 3 4]
CONF G101/S0 applies config 1 -> 2 ... adding=[0 1 2 3 4]
MIGR G101/S0 proposes installing shards=[0 1 2 3 4] config=2 keys=5 ...
MIGR G101/S0 installs shards=[0 1 2 3 4] config=2 keys=5 clients=1 ...
MIGR G100/S2 transfers shards=[0 1 2 3 4] config=2 keys=5
MIGR G100/S2 proposes deleting shards=[0 1 2 3 4] config=2 keys=5 ...
MIGR G100/S2 garbage-collects shards=[0 1 2 3 4] config=2 keys=5 ...
MIGR G101/S0 receives GC acknowledgement for shards=[0 1 2 3 4] config=2 keys=5
```

完整协议：

```text
ShardCtrler 产生配置 N
→ 各组将 Config N 写入自己的 Raft
→ 源组把转出的 shard 标记 MISSING
→ 目标组把新接收的 shard 标记 ADDING
→ 源组 Leader 收集对应 key/value 和去重表
→ 源组发送 PushShard RPC
→ 目标组 Leader 将 PushShardArgs 写入自己的 Raft
→ 目标组所有副本 Apply：安装数据、合并去重表、状态变为 WORKING
→ 目标组向源组发送 DeleteShard RPC
→ 源组 Leader 将 DeleteShardArgs 写入自己的 Raft
→ 源组所有副本删除旧数据，状态变为 WORKING
→ 源组返回 GC 成功
→ 两边都可以继续推进下一个配置
```

## 为什么迁移去重表

假设客户端的 Append 已在旧组执行，但回复丢失。随后 shard 被迁移，客户端向新组重试相同 ClientId/SeqNo。

如果只迁移键值，不迁移去重表，新组会重复执行 Append。因此 `PushShardArgs` 同时携带：

- shard 对应的键值数据。
- 各客户端已经执行的最大 SeqNo。

目标组合并去重表时对每个客户端取最大值。

## PushShard 和 DeleteShard 必须幂等

RPC 回复可能丢失，所以迁移请求会重复发送。

- 重复 PushShard 不能用旧数据覆盖已经更新的数据。
- 重复 DeleteShard 不能影响新的配置或新写入。
- 请求必须携带 Config Num，接收方据此判断请求是过期、未来还是当前配置。
- 真正修改数据前，Push/Delete 都先进入目标组自己的 Raft 日志。

`DUPL` 日志会标记已经安装或已经垃圾回收的迁移请求。

## ShardKV 快照必须包含什么

4B 快照包含：

```text
table        客户端去重表
data         键值数据
lastConfig   上一个配置
config       当前配置
shardsState  每个 shard 的 WORKING/MISSING/ADDING 状态
```

只保存数据是不够的。如果崩溃后丢失配置或 shard 状态，节点可能错误服务尚未迁入的 shard，或者忘记继续发送/清理迁移中的 shard。

## Challenge 的核心目标

### Challenge 1：及时删除旧 shard

目标组确认安装成功后，源组通过 Raft 提交 DeleteShard，删除已迁出的键值，防止旧数据长期占用快照空间。

### Challenge 2：不相关 shard 继续服务

某些 shard 正在迁移时，不受影响且处于 `WORKING` 的 shard 仍应继续处理请求。不能因为一个 shard 是 `ADDING/MISSING` 就暂停整个复制组。

## 4B 测试覆盖内容

| 测试 | 主要检查内容 |
|---|---|
| `TestStaticShards` | 静态多组分片与客户端路由 |
| `TestJoinLeave` | 组加入、离开和数据迁移 |
| `TestSnapshot` | 快照与配置迁移组合 |
| `TestMissChange` | 服务器错过中间配置后的逐配置追赶 |
| `TestConcurrent1/2/3` | 配置变化期间并发读写及重启 |
| `TestUnreliable1/2/3` | 不可靠网络下的迁移幂等性 |
| `TestChallenge1Delete` | 源组删除已迁出的数据 |
| `TestChallenge2Unaffected` | 不受迁移影响的 shard 保持可用 |
| `TestChallenge2Partial` | 部分 shard 完成迁移后及时恢复服务 |

## 筛选 4B 日志

只看配置推进：

```bash
rg ' CONF ' logs/lab4/4b-*.log
```

只看迁移握手：

```bash
rg ' MIGR ' logs/lab4/4b-*.log
```

跟踪某个复制组：

```bash
rg 'G100/' logs/lab4/4b-*.log
```

跟踪某批 shard：

```bash
rg 'shards=\[0 1 2 3 4\]' logs/lab4/4b-*.log
```

推荐阅读顺序：

```text
4A TestBasic4A
→ 4B TestStaticShards
→ 4B TestJoinLeave
→ 4B TestSnapshot
→ 4B TestMissChange
→ 并发/不可靠网络测试
→ Challenge 1/2
```

## 本次全量验证结果

- Lab 3A：14 个测试全部通过。
- Lab 3B：9 个测试全部通过。
- Lab 4A：2 个测试全部通过。
- Lab 4B：13 个测试全部通过，包括 Challenge 1 和 Challenge 2。
- Lab 3A、Lab 4A 和 Lab 4B 的代表性用例通过 race detector。
