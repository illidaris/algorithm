package fibonacci

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// todo爬楼梯 https://blog.csdn.net/a6661314/article/details/122720764
func FibonacciArr(n int) []int {
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			arr[i] = i
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
	return arr
}
