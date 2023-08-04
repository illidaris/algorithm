package recursive

import (
	"context"
	"errors"
	"fmt"
)

type Result struct {
	Data interface{}
	Err  error
}

func RecursiveFuncEmit(ctx context.Context, req ISetPage, maxNum uint, f func(ctx context.Context, req ISetPage) (interface{}, uint, error)) (<-chan Result, error) {
	outCh := make(chan Result, 10)
	go func() {
		defer func() {
			if rcvErr := recover(); rcvErr != nil {
				outCh <- Result{Err: errors.New(fmt.Sprintln("panic", rcvErr))}
				println(rcvErr)
			}
			if outCh != nil {
				close(outCh)
			}
		}()
		RecursiveFunc(func(b uint) uint {
			req.SetPage(b)
			data, total, err := f(ctx, req)
			outCh <- Result{Data: data, Err: err}
			return total
		}, maxNum)
	}()
	return outCh, nil
}

func RecursiveFunc(f func(uint) uint, maxNum uint) uint {
	return recursiveFunc(f, 1, 1, maxNum)
}

func recursiveFunc(f func(uint) uint, cursor, total, maxNum uint) uint {
	res := f(cursor)
	if cursor == 1 && res > 1 {
		total = res
	}
	if cursor > total-1 {
		return cursor
	}
	if cursor > maxNum-1 {
		return cursor
	}
	cursor++
	return recursiveFunc(f, cursor, total, maxNum)
}
