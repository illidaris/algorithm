package lottery

type ILottery interface {
	Invoke([]interface{}) []interface{}
}

type Base struct {
	Randm func(max int) int
}

// Draw draw by lottery func and pick some item
func DrawNum(source []interface{}, num uint, ls ...ILottery) []interface{} {
	if num == 0 {
		return nil
	}
	result := Draw(source, ls...)
	if maxNum := uint(len(result)); num > maxNum {
		num = maxNum
	}
	return result[:num]
}

// Draw draw by lottery func
func Draw(source []interface{}, ls ...ILottery) []interface{} {
	result := source
	for _, l := range ls {
		result = l.Invoke(source)
	}
	return result
}
