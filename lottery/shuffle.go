package lottery

var _ = ILottery(&Knuth{})

// Knuth knuth shuffie
type Knuth struct {
	Base
}

func (s *Knuth) Invoke(source []interface{}) []interface{} {
	for i := len(source) - 1; i > 0; i-- {
		r := s.Randm(i + 1)
		source[i], source[r] = source[r], source[i]
	}
	return source
}
