package lottery

type ILottery interface {
	Invoke([]interface{}) []interface{}
}

type Base struct {
	Randm func(max int) int
}

// Draw draw by lottery func and pick some item
func Draw(source []interface{}, num uint, ls ...ILottery) []interface{} {
	if num == 0 {
		return nil
	}
	result := source
	for _, l := range ls {
		result = l.Invoke(source)
	}
	if max := uint(len(result)); num > max {
		num = max
	}
	return result[:num]
}
