package distinct

type None struct{}

var none None

type Set struct {
	itemMap map[interface{}]IRecord
	items   []IRecord
}

func NewSet() *Set {
	return &Set{
		itemMap: map[interface{}]IRecord{},
		items:   []IRecord{},
	}
}

func (s *Set) Add(args ...IRecord) {
	for _, arg := range args {
		id := arg.GetID()
		if _, ok := s.itemMap[id]; !ok {
			s.itemMap[id] = arg
			s.items = append(s.items, arg)
		}
	}
}

func (s *Set) Slice() []IRecord {
	res := s.items
	return res
}

// 去重

func (s *Set) Invoke(args []IRecord) []IRecord {

	// result := []int{}
	// m := make(map[int]bool) //map的值不重要
	// for _, v := range s {
	// 	if _, ok := m[v]; !ok {
	// 		result = append(result, v)
	// 		m[v] = true
	// 	}
	// }

	// return result
	return nil

}

// 删除指定元素
// DeleteSlice4 删除指定元素。
// func DeleteSlice4(a []int, elem int) []int {
//     tgt := a[:0]
//     for _, v := range a {
//         if v != elem {
//             tgt = append(tgt, v)
//         }
//     }
//     return tgt
// }
// 性能最佳
// func DeleteSlice3(a []int, elem int) []int {
//     j := 0
//     for _, v := range a {
//         if v != elem {
//             a[j] = v
//             j++
//         }
//     }
//     return a[:j]
// }
