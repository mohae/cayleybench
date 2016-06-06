// Copyright 2016 Markus W Mahlberg <markus.mahlberg@me.com>
//
// This file is part of cayleybench.
//
// cayleybench is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cayleybench is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with cayleybench.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/cayley"
	"github.com/google/cayley/graph"
	_ "github.com/google/cayley/graph/bolt"
	_ "github.com/google/cayley/graph/mongo"
	_ "github.com/google/cayley/graph/sql"
	"github.com/mohae/benchutil"
	cb "github.com/mwmahlberg/cayleybench"
)

// flags
var (
	output         string
	format         string
	murl           string
	pgurl          string
	cpus           string
	section        bool
	sectionHeaders bool
	nameSections   bool
	systemInfo     bool
	sleep          int64
	once           sync.Once
)

func init() {
	flag.BoolVar(&nameSections, "namesections", false, "use group as section name: some restrictions apply")
	flag.BoolVar(&nameSections, "n", false, "use group as section name: some restrictions apply")
	flag.BoolVar(&section, "sections", false, "don't separate groups of tests into sections")
	flag.BoolVar(&section, "s", false, "don't separate groups of tests into sections")
	flag.BoolVar(&sectionHeaders, "sectionheader", false, "if there are sections, add a section header row")
	flag.BoolVar(&sectionHeaders, "h", false, "if there are sections, add a section header row")
	flag.BoolVar(&systemInfo, "sysinfo", false, "add the system information to the output")
	flag.BoolVar(&systemInfo, "i", false, "add the system information to the output")
	flag.StringVar(&pgurl, "p", "postgres://postgres@postgres/postgres?sslmode=disable&connect_timeout=0", "PostGres url to connect to")
	flag.Int64Var(&sleep, "sleep", 0, "waits this number of seconds before executing the tests")
	flag.Int64Var(&sleep, "z", 0, "waits this number of seconds before executing the tests")
	flag.StringVar(&cpus, "cpus", "", "comma-separated list of number of CPUs to use for each test")
	flag.StringVar(&cpus, "c", "", "comma-separated list of number of CPUs to use for each test")
	flag.StringVar(&format, "format", "txt", "format of output")
	flag.StringVar(&format, "f", "txt", "format of output")
	flag.StringVar(&murl, "mongo", "mongodb:27017", "MongoDB url to connect to")
	flag.StringVar(&murl, "m", "mongodb:27017", "MongoDB url to connect to")
	flag.StringVar(&output, "o", "stdout", "output destination (short)")
	flag.StringVar(&pgurl, "pg", "postgres://postgres@postgres/postgres?sslmode=disable&connect_timeout=0", "PostGres url to connect to")
	flag.StringVar(&pgurl, "p", "postgres://postgres@postgres/postgres?sslmode=disable&connect_timeout=0", "PostGres url to connect to")
	flag.StringVar(&output, "output", "stdout", "output destination")
}

func main() {
	flag.Parse()
	if sleep > 0 {
		time.Sleep(time.Duration(sleep) * time.Second)
	}
	CPUs := parseCPUs()

	// if there weren't any CPUs specified, set it to the maxprocs
	if len(CPUs) == 0 {
		CPUs = append(CPUs, runtime.NumCPU())
	}

	done := make(chan struct{})
	// start the visual ticker
	go benchutil.Dot(done)
	// set the output
	var w io.Writer
	var err error

	switch output {
	case "stdout":
		w = os.Stdout
	default:
		w, err = os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer w.(*os.File).Close()
	}

	// get the benchmark for the desired format
	// process the output
	var bench benchutil.Benchmarker
	switch format {
	case "csv":
		bench = benchutil.NewCSVBench(w)
	case "md":
		bench = benchutil.NewMDBench(w)
		bench.(*benchutil.MDBench).GroupAsSectionName = nameSections
	default:
		bench = benchutil.NewStringBench(w)
	}
	bench.SectionPerGroup(section)
	bench.SectionHeaders(sectionHeaders)
	bench.IncludeSystemInfo(systemInfo)

	// set the header info
	bench.SetGroupColumnHeader("CPUs")
	bench.SetNameColumnHeader("Database")

	benchAll(bench, CPUs)
}

func benchAll(bench benchutil.Benchmarker, CPUs []int) {

	for _, cpu := range CPUs {
		cpu = runtime.GOMAXPROCS(cpu)
		b := benchutil.NewBench("bolt insert")
		b.Group = strconv.Itoa(cpu)
		b.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchBoltInsert))
		bench.Append(b)

		b = benchutil.NewBench("mem insert")
		b.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchMemInsert))
		bench.Append(b)

		b = benchutil.NewBench("mongo insert")
		b.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchMongoInsert))
		bench.Append(b)

		b = benchutil.NewBench("postgress insert")
		b.Result = benchutil.ResultFromBenchmarkResult(testing.Benchmark(BenchPostgresInsert))
		bench.Append(b)
	}
}

func BenchMemInsert(b *testing.B) {
	var (
		err  error
		memg *cayley.Handle
		memw *cb.QuadWriter
	)

	if memg, err = cayley.NewMemoryGraph(); err != nil {
		b.Fatal("mem: Can not create MemoryGraph", err)
	}
	defer memg.Close()

	memw = cb.NewQuadWriter(memg)
	memw.InitQuads(b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		memw.WriteQuad()
	}
}

func BenchPostgresInsert(b *testing.B) {

	var (
		err error
		pg  *cayley.Handle
		pw  *cb.QuadWriter
	)

	once.Do(
		func() {
			if err = graph.InitQuadStore("sql", pgurl, nil); err != nil {
				b.Fatal("pgsql: Can not init QuadStore", err)
			}
		})

	if pg, err = cayley.NewGraph("sql", pgurl, nil); err != nil {
		b.Fatal("pgsql: Can not create QuadStore", err)
	}

	defer pg.Close()

	pw = cb.NewQuadWriter(pg)
	pw.InitQuads(b.N)

	b.ResetTimer()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				pw.WriteQuad()
			}
		})
	runtime.GC()
}

func BenchMongoInsert(b *testing.B) {

	var (
		err error
		mg  *cayley.Handle
		mw  *cb.QuadWriter
	)

	if err = graph.InitQuadStore("mongo", murl, nil); err != nil {
		b.Fatal("mongo: Can not init QuadStore", err)
	}

	mg, err = cayley.NewGraph("mongo", murl, nil)
	if err != nil {
		b.Fatal("mongo: Can not create QuadStore", err)
	}
	defer mg.Close()

	mw = cb.NewQuadWriter(mg)
	mw.InitQuads(b.N)

	b.ResetTimer()

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				mw.WriteQuad()
			}
		})

}

func BenchBoltInsert(b *testing.B) {

	var (
		err error
		bg  *cayley.Handle
		bw  *cb.QuadWriter
	)

	f, err := ioutil.TempFile("", "cayleybench-bolt")
	if err != nil {
		b.Fatal("boltdb: Error creating temporary file for BoltDB", err)
	}

	fn := f.Name()
	f.Close()

	defer os.Remove(fn)

	if err = graph.InitQuadStore("bolt", fn, nil); err != nil {
		b.Fatal("boltdb: Can not init QuadStore", err)
	}

	if bg, err = cayley.NewGraph("bolt", fn, nil); err != nil {
		b.Fatal("boltdb: Can not create QuadStore", err)
	}
	defer bg.Close()

	bw = cb.NewQuadWriter(bg)

	bw.InitQuads(b.N)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bw.WriteQuad()
		}
	})
}

func parseCPUs() []int {
	var CPUs []int
	vals := strings.Split(cpus, ",")
	for _, val := range vals {
		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}
		cpu, err := strconv.Atoi(val)
		if err != nil || cpu <= 0 {
			fmt.Fprintf(os.Stderr, "invalid cpu value: %s\n", val)
			os.Exit(1)
		}
		CPUs = append(CPUs, cpu)
	}
	return CPUs
}
