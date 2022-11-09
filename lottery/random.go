package lottery

var _ = ILottery(&Random{})

// Randmon random token
type Random struct {
	Base
}

func (s *Random) Invoke(source []interface{}) []interface{} {
	if source == nil {
		return nil
	}
	max := len(source)
	tokens := []interface{}{}
	temp := source
	for i := 0; i < max; i++ {
		subLen := len(temp)
		if subLen == 0 {
			break
		}
		index := s.Randm(subLen)
		tokens = append(tokens, temp[index])
		temp = DeleteSlice(temp, index)
	}
	return tokens
}

func DeleteSlice(s []interface{}, index int) []interface{} {
	j := 0
	for i, v := range s {
		if i != index {
			s[j] = v
			j++
		}
	}
	return s[:j]
}
