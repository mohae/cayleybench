# cayleybench

I wrote cayleybench because the question arose on Cayley's slack channel
which option would be prefereable as a backend, BoltDB or MongoDB.

## Install

    go get github.com/mwmahlberg/cayleybench

## Run
    $ cd $GOPATH/src/github.com/mwmahlberg/cayleybench
    $ go test -run=XXX -bench=. -benchmem -cpu 2,4,6,8
    testing: warning: no tests to run
    PASS
    BenchmarkMongoInsert-2	    1000	   1370244 ns/op	   32592 B/op	     523 allocs/op
    BenchmarkMongoInsert-4	    1000	   1286237 ns/op	   32606 B/op	     523 allocs/op
    BenchmarkMongoInsert-6	    1000	   1395686 ns/op	   32628 B/op	     524 allocs/op
    BenchmarkMongoInsert-8	    1000	   1440273 ns/op	   32603 B/op	     523 allocs/op
    BenchmarkBoltInserts-2	    1000	   2119516 ns/op	  121353 B/op	     267 allocs/op
    BenchmarkBoltInserts-4	     500	   2533091 ns/op	  145026 B/op	     290 allocs/op
    BenchmarkBoltInserts-6	     500	   2724483 ns/op	  155769 B/op	     305 allocs/op
    BenchmarkBoltInserts-8	     500	   2873391 ns/op	  159696 B/op	     314 allocs/op
    ok  	github.com/mwmahlberg/cayleybench	14.942s
