package lottery

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestDraw(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	convey.Convey("draw", t, func() {
		ids := []interface{}{}
		total := rand.Intn(100) + 1
		for i := 0; i < total; i++ {
			ids = append(ids, i)
		}

		convey.Convey("no random just pick", func() {
			k := Random{}
			k.Randm = func(max int) int {
				return rand.Intn(max)
			}

			convey.Convey("no random pick 0", func() {
				result := Draw(ids, 0)
				convey.So(result, convey.ShouldBeNil)
			})

			convey.Convey("no random pick less then source len ", func() {
				result := Draw(ids, uint(total-1))
				bs, _ := json.Marshal(result)
				fmt.Printf("%s", string(bs))
				convey.So(len(result), convey.ShouldEqual, total-1)
				for index, v := range result {
					convey.So(v, convey.ShouldEqual, ids[index])
				}
			})

			convey.Convey("no random pick more then source len", func() {
				result := Draw(ids, uint(total+1))
				bs, _ := json.Marshal(result)
				fmt.Printf("%s", string(bs))
				convey.So(len(result), convey.ShouldEqual, total)
				for index, v := range result {
					convey.So(v, convey.ShouldEqual, ids[index])
				}
			})
		})

		convey.Convey("with randmon", func() {
			k := &Knuth{}
			k.Randm = func(max int) int {
				return rand.Intn(max)
			}
			r := &Random{}
			r.Randm = func(max int) int {
				return rand.Intn(max)
			}

			examples := map[string][]ILottery{"knuth": {k}, "random": {r}, "mixed": {k, r}}
			for key, example := range examples {
				convey.Convey(fmt.Sprintf("%s pick 0", key), func() {
					result := Draw(ids, 0, example...)
					convey.So(result, convey.ShouldBeNil)
				})

				convey.Convey(fmt.Sprintf("%s pick less then source len", key), func() {
					result := Draw(ids, uint(total-1), example...)
					bs, _ := json.Marshal(result)
					fmt.Printf("%s", string(bs))
					convey.So(len(result), convey.ShouldEqual, total-1)
				})

				convey.Convey(fmt.Sprintf("%s pick more then source len", key), func() {
					result := Draw(ids, uint(total+1), example...)
					bs, _ := json.Marshal(result)
					fmt.Printf("%s", string(bs))
					convey.So(len(result), convey.ShouldEqual, total)
				})
			}

		})
	})
}

func BenchmarkDraw(b *testing.B) {
	ids := []interface{}{}
	total := 1000000
	k := &Knuth{}
	k.Randm = func(max int) int {
		return rand.Intn(max)
	}
	r := &Random{}
	r.Randm = func(max int) int {
		return rand.Intn(max)
	}
	for i := 0; i < total; i++ {
		ids = append(ids, i)
	}
	for n := 0; n < b.N; n++ {
		result := Draw(ids, 1000, k)
		if len(result) < 1000 {
			b.Error("less then 1000")
		}
	}
}
