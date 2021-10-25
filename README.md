# Binary Search Tree
Implements search for the deepest node inside the btree

To run the challenge example:
```bash
go run cmd/main/main.go
```
Output (draws the graph and returns the result):
```bash
              ----[7
                     ----[9
       ----[11
----[12
              ----[82
       ----[90
deepest 9 depth 3
```

## Test 
```bash
go test
```

```bash
PASS
ok  	github.com/russiansmack/btree	17.179s
```

## Benchmarks 
```bash
go test -bench=. -benchtime=10000000x
```

```bash
goos: linux
goarch: amd64
pkg: github.com/russiansmack/btree
BenchmarkNode_Insert1-16              	10000000	         5.42 ns/op
BenchmarkNode_Insert5-16              	10000000	        13.1 ns/op
BenchmarkNode_Insert10-16             	10000000	        17.7 ns/op
BenchmarkNode_Insert100-16            	10000000	        37.3 ns/op
BenchmarkNode_Insert100000-16         	10000000	       309 ns/op
BenchmarkNode_FindDeepestSimple-16    	10000000	         0.000189 ns/op
BenchmarkNode_FindDeepest1-16         	10000000	         0.000140 ns/op
BenchmarkNode_FindDeepest5-16         	10000000	         0.000189 ns/op
BenchmarkNode_FindDeepest10-16        	10000000	         0.000244 ns/op
BenchmarkNode_FindDeepest100-16       	10000000	         0.000328 ns/op
BenchmarkNode_FindDeepest100000-16    	10000000	         0.254 ns/op
PASS
ok  	github.com/russiansmack/btree	27.877s
```
