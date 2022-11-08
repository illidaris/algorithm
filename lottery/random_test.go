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
	convey.Convey("knuth run", t, func() {
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
