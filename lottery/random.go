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
		switch {
		case index == 0:
			tokens = append(tokens, temp[0])
			temp = temp[1:]
		case index+1 == subLen:
			tokens = append(tokens, temp[index])
			temp = temp[:index]
		default:
			tokens = append(tokens, temp[index-1])
			temp = append(temp[:index-1], temp[index:]...)
		}
	}
	return tokens
}
