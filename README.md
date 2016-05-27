# cayleybench

I wrote cayleybench because the question arose on Cayley's slack channel
which option would be prefereable as a backend, BoltDB or MongoDB.

## Install

    go get github.com/mwmahlberg/cayleybench

There will be a docker image, too.    

## Run

### In-Memory Graph

    $ go test -run=XXX -bench=.*Mem -benchmem -cpu 2,4,6,8
    testing: warning: no tests to run
    PASS
    BenchmarkMemInsert-2	  100000	     20491 ns/op	    3590 B/op	       5 allocs/op
    BenchmarkMemInsert-4	  100000	     16146 ns/op	    3591 B/op	       5 allocs/op
    BenchmarkMemInsert-6	  100000	     18554 ns/op	    3590 B/op	       5 allocs/op
    BenchmarkMemInsert-8	  100000	     22909 ns/op	    3591 B/op	       5 allocs/op
    ok  	github.com/mwmahlberg/cayleybench	11.632s

### MongoDB Graph    

    $ go test -run=XXX -bench=.*Mongo -benchmem -cpu 2,4,6,8
    testing: warning: no tests to run
    PASS
    BenchmarkMongoInsert-2	    1000	   1487063 ns/op	   32591 B/op	     523 allocs/op
    BenchmarkMongoInsert-4	    1000	   1368688 ns/op	   32603 B/op	     523 allocs/op
    BenchmarkMongoInsert-6	    1000	   1467024 ns/op	   32630 B/op	     524 allocs/op
    BenchmarkMongoInsert-8	    1000	   1467883 ns/op	   32608 B/op	     523 allocs/op
    ok  	github.com/mwmahlberg/cayleybench	46.511s

### BoltDB Graph

    $ go test -run=XXX -bench=.*Bolt -benchmem -cpu 2,4,6,8
    testing: warning: no tests to run
    PASS
    BenchmarkBoltInserts-2	    1000	   2054540 ns/op	  117495 B/op	     265 allocs/op
    BenchmarkBoltInserts-4	    1000	   2052377 ns/op	  117890 B/op	     265 allocs/op
    BenchmarkBoltInserts-6	    1000	   2140950 ns/op	  117535 B/op	     265 allocs/op
    BenchmarkBoltInserts-8	    1000	   2098326 ns/op	  117848 B/op	     265 allocs/op
    ok  	github.com/mwmahlberg/cayleybench	10.265s

### PostgreSQL

> **Please Note** This is only added for reference purposes
> I only have run this against a PostgreSQL docker image.
> **DO NOT TAKE THESE VALUES SERIOUS**

    $ go test -run=XXX -bench=.*Post -benchmem -cpu 2,4,6,8
    testing: warning: no tests to run
    PASS
    BenchmarkPostgresInsert-2	    1000	   2027693 ns/op	   69503 B/op	      84 allocs/op
    BenchmarkPostgresInsert-4	    1000	   1926262 ns/op	   69542 B/op	      84 allocs/op
    BenchmarkPostgresInsert-6	     500	   2002072 ns/op	   69687 B/op	      86 allocs/op
    BenchmarkPostgresInsert-8	     500	   2199349 ns/op	   69797 B/op	      87 allocs/op
    ok  	github.com/mwmahlberg/cayleybench	9.109s
