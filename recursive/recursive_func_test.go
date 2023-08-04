package recursive

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestRecursiveFuncEmit(t *testing.T) {
	ctx := context.Background()

	ch, err := RecursiveFuncEmit(ctx, []string{"序号", "字段1", "字段2"}, &Bag{ID: 1}, 0, func(ctx context.Context, req ISetPage) ([][]string, uint) {
		result := [][]string{{"1", "xxxxxxxxx", "ssss"}, {"2", "xxxxxxxxx", "ssss"}}
		return result, 3
	})
	if err != nil {
		t.Error(err)
	}
	for row := range ch {
		println(strings.Join(row, ","))
	}

}

func TestCh(t *testing.T) {
	ch := make(chan *Bag, 10)
	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- &Bag{ID: uint(i)}
			println("<-", i)
		}
		println("end")
	}()
	for v := range ch {
		time.Sleep(time.Millisecond * 10)
		println("->", v.ID)
	}

}

func TestRecursion(t *testing.T) {
	ch := make(chan string, 1)
	bag := &Bag{ID: 1}
	go func() {
		defer close(ch)
		RecursiveFunc(func(b uint) uint {
			bag.ID = b
			ch <- fmt.Sprintf("数据%v", bag)
			return 5
		}, 2)
	}()
	for v := range ch {
		time.Sleep(time.Millisecond * 10)
		println("->", v)
	}
}

type Bag struct {
	ID uint
}

func (e *Bag) SetPage(i uint) {
	e.ID = i
}

func (e *Bag) GetPageIndex() int {
	return int(e.ID)
}
