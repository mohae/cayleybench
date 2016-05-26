package mongodb

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/google/cayley"
	"github.com/google/cayley/graph"
	_ "github.com/google/cayley/graph/bolt"
	_ "github.com/google/cayley/graph/mongo"

	"github.com/pborman/uuid"
)

var (
	cm, cb *cayley.Handle
)

type UUIDs struct {
	sync.Mutex
	values  []string
	counter int32
}

func (u *UUIDs) GetUUID() (id string) {

	u.Lock()
	defer u.Unlock()

	id = u.values[u.counter]
	u.counter++
	return
}

func NewUUIDs(count int) (ids *UUIDs) {
	ids = &UUIDs{}
	ids.values = make([]string, count, count)

	for i := 0; i < count; i++ {
		ids.values[i] = uuid.NewRandom().String()
	}

	return
}

func TestMain(m *testing.M) {
	var err error
	graph.InitQuadStore("mongo", "mongodb://127.0.0.1:27017", nil)

	cm, err = cayley.NewGraph("mongo", "mongodb://127.0.0.1:27017", nil)

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to MongoDB")
	}

	f, err := ioutil.TempFile("", "cayleybench-bolt")

	if err != nil {
		panic("Error creating temporary file for BoltDB")
	}

	fn := f.Name()
	f.Close()

	defer os.Remove(fn)

	if err = graph.InitQuadStore("bolt", fn, nil); err != nil {
		panic(err)
	}

	cb, err = cayley.NewGraph("bolt", fn, nil)

	if err != nil {
		panic(err)
	}

	m.Run()
}

func BenchmarkMongoInsert(b *testing.B) {
	uuids := NewUUIDs(b.N*2 + 1)
	b.ResetTimer()
	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				cm.AddQuad(cayley.Quad(uuids.GetUUID(), "follows", uuids.GetUUID(), ""))
			}
		})
}

func BenchmarkBoltInserts(b *testing.B) {
	uuids := NewUUIDs(b.N*2 + 1)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			cb.AddQuad(cayley.Quad(uuids.GetUUID(), "follows", uuids.GetUUID(), ""))
		}
	})
}
