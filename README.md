[![Build Status](https://travis-ci.org/mwmahlberg/cayleybench.svg?branch=master)](https://travis-ci.org/mwmahlberg/cayleybench)
[![Docker Pulls](https://img.shields.io/docker/pulls/mwmahlberg/cayleybench.svg)][hub]
[![](https://imagelayers.io/badge/mwmahlberg/cayleybench:latest.svg)](https://imagelayers.io/?images=mwmahlberg/cayleybench:latest 'Get your own badge on imagelayers.io')

# cayleybench

I wrote cayleybench because the question arose on Cayley's slack channel
which option would be prefereable as a backend, BoltDB or MongoDB.

## Install

    go get github.com/mwmahlberg/cayleybench

## Run as docker-compose

    $ curl -o docker-compose.yml https://raw.githubusercontent.com/mwmahlberg/cayleybench/master/docker-compose.yml
    $ docker-compose up --abort-on-container-exit
    Creating network "cayleybench_default" with the default driver
    Creating cayleybench_postgres_1
    Creating cayleybench_mongodb_1
    Creating benchrunner
    Attaching to cayleybench_postgres_1, cayleybench_mongodb_1, benchrunner
    postgres_1   | WARNING: no logs are available with the 'none' log driver
    mongodb_1    | WARNING: no logs are available with the 'none' log driver
    benchrunner  | testing: warning: no tests to run
    benchrunner  | PASS
    benchrunner  | BenchmarkMemInsert       	   20000	     50126 ns/op	    3672 B/op	       5 allocs/op
    benchrunner  | BenchmarkMemInsert-2     	   30000	     46544 ns/op	    3871 B/op	       5 allocs/op
    benchrunner  | BenchmarkMemInsert-4     	   50000	     30957 ns/op	    3596 B/op	       5 allocs/op
    benchrunner  | BenchmarkMemInsert-6     	  100000	     46518 ns/op	    3591 B/op	       5 allocs/op
    benchrunner  | BenchmarkMemInsert-8     	   50000	     23543 ns/op	    3597 B/op	       5 allocs/op
    benchrunner  | BenchmarkPostgresInsert  	    1000	   1196971 ns/op	   69422 B/op	      84 allocs/op
    benchrunner  | BenchmarkPostgresInsert-2	    1000	   1261531 ns/op	   69486 B/op	      84 allocs/op
    benchrunner  | BenchmarkPostgresInsert-4	    1000	   1393058 ns/op	   69523 B/op	      85 allocs/op
    benchrunner  | BenchmarkPostgresInsert-6	    1000	   1830245 ns/op	   69614 B/op	      85 allocs/op
    benchrunner  | BenchmarkPostgresInsert-8	     300	   3543214 ns/op	   70149 B/op	      90 allocs/op
    benchrunner  | BenchmarkMongoInsert     	    1000	   1684434 ns/op	   34364 B/op	     554 allocs/op
    benchrunner  | BenchmarkMongoInsert-2   	    1000	   1711045 ns/op	   34383 B/op	     554 allocs/op
    benchrunner  | BenchmarkMongoInsert-4   	    1000	   1735010 ns/op	   34409 B/op	     554 allocs/op
    benchrunner  | BenchmarkMongoInsert-6   	    1000	   1779852 ns/op	   34414 B/op	     554 allocs/op
    benchrunner  | BenchmarkMongoInsert-8   	    1000	   1835134 ns/op	   34438 B/op	     555 allocs/op
    benchrunner  | BenchmarkBoltInserts     	    1000	   2339768 ns/op	   62924 B/op	     267 allocs/op
    benchrunner  | BenchmarkBoltInserts-2   	    1000	   2131284 ns/op	   63401 B/op	     267 allocs/op
    benchrunner  | BenchmarkBoltInserts-4   	    1000	   2163644 ns/op	   63072 B/op	     267 allocs/op
    benchrunner  | BenchmarkBoltInserts-6   	    1000	   2205209 ns/op	   63450 B/op	     267 allocs/op
    benchrunner  | BenchmarkBoltInserts-8   	    1000	   2131469 ns/op	   62792 B/op	     267 allocs/op
    benchrunner exited with code 0
    Aborting on container exit...
    Stopping cayleybench_mongodb_1 ... done
    Stopping cayleybench_postgres_1 ... done

## Run manually

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
