package lottery

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestRandom_Invoke(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	convey.Convey("random run", t, func() {
		ids := []interface{}{}
		for i := 0; i < 10; i++ {
			ids = append(ids, i)
		}

		convey.Convey("run 10 success", func() {
			k := Random{}
			k.Randm = func(max int) int {
				return rand.Intn(max)
			}
			result := k.Invoke(ids)
			bs, _ := json.Marshal(result)
			fmt.Printf("%s", string(bs))
		})

		convey.Convey("run empty source", func() {
			k := Random{}
			k.Randm = func(max int) int {
				return rand.Intn(max)
			}
			result := k.Invoke([]interface{}{})
			bs, _ := json.Marshal(result)
			fmt.Printf("%s", string(bs))
		})

		convey.Convey("run nil source", func() {
			k := Random{}
			k.Randm = func(max int) int {
				return rand.Intn(max)
			}
			result := k.Invoke(nil)
			bs, _ := json.Marshal(result)
			fmt.Printf("%s", string(bs))
		})
	})
}

func TestDeleteSlice(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	convey.Convey("delete slice", t, func() {
		raw := []interface{}{}
		for i := 0; i < 10; i++ {
			raw = append(raw, i)
		}

		convey.Convey("remove first element", func() {
			ids := raw
			ids = DeleteSlice(ids, 0)
			convey.So(len(ids), convey.ShouldEqual, 9)
			convey.So(ids[0], convey.ShouldEqual, 1)
			convey.So(ids[len(ids)-1], convey.ShouldEqual, 9)
			bs, _ := json.Marshal(ids)
			fmt.Printf("%s", string(bs))
		})

		convey.Convey("remove last element", func() {
			ids := raw
			ids = DeleteSlice(ids, len(ids)-1)
			convey.So(len(ids), convey.ShouldEqual, 9)
			convey.So(ids[0], convey.ShouldEqual, 0)
			convey.So(ids[len(ids)-1], convey.ShouldEqual, 8)
			bs, _ := json.Marshal(ids)
			fmt.Printf("%s", string(bs))
		})

		convey.Convey("remove middle element", func() {
			ids := raw
			ids = DeleteSlice(ids, 1)
			convey.So(len(ids), convey.ShouldEqual, 9)
			convey.So(ids[0], convey.ShouldEqual, 0)
			convey.So(ids[len(ids)-1], convey.ShouldEqual, 9)
			bs, _ := json.Marshal(ids)
			fmt.Printf("%s", string(bs))
		})
	})
}

func getSlice(num int) []interface{} {
	ids := []interface{}{}
	for i := 0; i < num; i++ {
		ids = append(ids, i)
	}
	return ids
}

func BenchmarkDeleteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DeleteSlice(getSlice(10000), 10)
	}
}
