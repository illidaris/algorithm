package recursive

import "context"

func RecursiveFuncEmit(ctx context.Context, title []string, req ISetPage, f func(ctx context.Context, req ISetPage) ([][]string, uint)) (<-chan []string, error) {
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
		})
	}()
	return outCh, nil
}

func RecursiveFunc(f func(uint) uint) uint {
	return recursiveFunc(f, 1, 0)
}

func recursiveFunc(f func(uint) uint, begin, total uint) uint {
	res := f(begin)
	if begin == 1 && total == 0 {
		total = res
	}
	if begin > total-1 {
		return 0
	}
	begin++
	return recursiveFunc(f, begin, total)
}
