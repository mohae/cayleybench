
[![LGPLv3](https://img.shields.io/badge/License-LGPLv3-337128.svg?style=flat-square)](http://www.gnu.org/licenses/lgpl-3.0-standalone.html)
[![Build Status](https://img.shields.io/travis/mwmahlberg/cayleybench/master.svg?style=flat-square)](https://travis-ci.org/mwmahlberg/cayleybench) [![Godoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/mwmahlberg/cayleybench)
[![Docker Pulls](https://img.shields.io/docker/pulls/mwmahlberg/cayleybench.svg?style=flat-square)](https://hub.docker.com/r/mwmahlberg/cayleybench/)
[![Image Layers](https://imagelayers.io/badge/mwmahlberg/cayleybench:latest.svg)](https://imagelayers.io/?images=mwmahlberg/cayleybench:latest 'Get your own badge on imagelayers.io')

# cayleybench

This project provides an easy to use benchmark for the [Cayley graph database][cayley].

I wrote cayleybench because the question arose on Cayley's slack channel([#cayley][cayleyslack]) which option would be prefereable as a backend, [BoltDB][bolt] or [MongoDB][mongo]. Because it was easy, I added a benchmark for the In-Memory Graph. Adding one for [PostgreSQL][pgsql] was only the next logical step.

I noticed very quickly that setting up a MongoDB and a PostgreSQL server might be a bit complicated for some,
so I decided  that it would be best if cayleybench could be run with [docker-compose][compose].

# Run

## With `docker-compose`

If you just want to run the benchmark and do not have MongoDB and/or PostgreSQL installed,
you should use this option.

It downloads a known good MongoDB and PostgreSQL docker image as well as the latest version of
`cayleybench`, fires up the machines and runs the benchmarks against those docker images.

To start, simply type

    $ curl -o docker-compose.yml https://raw.githubusercontent.com/mwmahlberg/cayleybench/master/docker-compose.yml
    $ docker-compose up --abort-on-container-exit

After the benchmark is done, you might want to remove the stopped containers with

    $ docker-compose down

## With `docker`

If you want to test against your existing machines, you can run the `cayleybench` docker image on it's own
by typing

    $ docker run -it mwmahlberg/cayleybench -test.bench=.\*(Mongo|Mem|Bolt|Postgres) -mongo=<mongoUrl> -pg=<pgURL>

Of course, you can leave out any test you wish by removing it from the regular expression passed to `-test.bench`.
If you skip the MongoDB benchmark, you do not need to specify `-mongo`. The same goes for the PostgreSQL benchmark and `-pg`,
respectively.

For the details on `<mongoUrl>` please [see 'Connection String URI Format' in the MongoDB documentation][mongo:uri].

`cayleybench` uses [github.com/lib/pq][pq] for accessing PostgreSQL. Please [see the godoc][pq:godoc] for format of `<pgURL>`.

### Command line parameters

You can pass any of the following parameters after `docker run -it mwmahlberg/cayleybench`

    -alsologtostderr
        log to standard error as well as files
    -ignoredup
        Don't stop loading on duplicated key on add
    -ignoremissing
        Don't stop loading on missing key on delete
    -log_backtrace_at value
        when logging hits line file:N, emit a stack trace (default :0)
    -log_dir string
        If non-empty, write log files in this directory
    -logstashtype string
        enable logstash logging and define the type
    -logstashurl string
        logstash url and port (default "172.17.42.1:5042")
    -logtostderr
        log to standard error instead of files
    -mongo string
        MongoDB url to connect to (default "mongodb:27017")
    -pg string
        PostGres url to connect to (default "postgres://postgres@postgres/postgres?sslmode=disable&connect_timeout=0")
    -sleep int
        waits this number of seconds before executing the tests
    -stderrthreshold value
        logs at or above this threshold go to stderr
    -test.bench string
        regular expression to select benchmarks to run
    -test.benchmem
        print memory allocations for benchmarks
    -test.benchtime duration
        approximate run time for each benchmark (default 1s)
    -test.blockprofile string
        write a goroutine blocking profile to the named file after execution
    -test.blockprofilerate int
        if >= 0, calls runtime.SetBlockProfileRate() (default 1)
    -test.count n
        run tests and benchmarks n times (default 1)
    -test.coverprofile string
        write a coverage profile to the named file after execution
    -test.cpu string
        comma-separated list of number of CPUs to use for each test
    -test.cpuprofile string
        write a cpu profile to the named file during execution
    -test.memprofile string
        write a memory profile to the named file after execution
    -test.memprofilerate int
        if >=0, sets runtime.MemProfileRate
    -test.outputdir string
        directory in which to write profiles
    -test.parallel int
        maximum test parallelism (default 2)
    -test.run string
        regular expression to select tests and examples to run
    -test.short
        run smaller test suite to save time
    -test.timeout duration
        if positive, sets an aggregate time limit for all tests
    -test.trace string
        write an execution trace to the named file after execution
    -test.v
        verbose: print additional output
    -v value
        log level for V logs
    -vmodule value
        comma-separated list of pattern=N settings for file-filtered logging

Please note that the profiling output would be written to the docker image, so you would need to [define a volume][docker:data].

[cayley]: http://cayley.io "Cayley project page"
[cayleyslack]: https://gophers.slack.com/messages/cayley/ "Cayley Slack channel"
[compose]: https://www.docker.com/products/docker-compose#/overview "docker-compose product overview"
[bolt]:  https://github.com/boltdb/bolt#bolt--- "README of BoltDB"
[mongo]: https://docs.mongodb.com/manual/ "MongoDB documentation"
[mongo:uri]: https://docs.mongodb.com/manual/reference/connection-string/ "Connection String URI Format"
[pgsql]: https://www.postgresql.org "PostgreSQL project page"
[docker:data]: https://docs.docker.com/engine/userguide/containers/dockervolumes/ "Docker Docs: Manage data in containers"
[pq]: https://github.com/lib/pq "Project Page of lib/pq"
[pq:godoc]: https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters "Godoc of lib/pq: Connection String Parameters"
