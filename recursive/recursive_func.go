package recursive

import "context"

func RecursiveFuncEmit(ctx context.Context, title []string, req ISetPage, maxNum uint, f func(ctx context.Context, req ISetPage) ([][]string, uint)) (<-chan []string, error) {
	outCh := make(chan []string, 10)
	go func() {
		defer close(outCh)
		if len(title) == 0 {
			outCh <- title
		}
		RecursiveFunc(func(b uint) uint {
			req.SetPage(b)
			rows, total := f(ctx, req)
			for _, row := range rows {
				outCh <- row
			}
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
