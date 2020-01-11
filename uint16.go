package slice

import (
	"sort"
)

// UInter16 is the interface that handles a uint16 collection.
type UInter16 interface {
	Append(...uint16) UInter16
	Bounds(int) bool
	Concatenate(UInter16) UInter16
	Each(func(int, uint16)) UInter16
	EachBreak(func(int, uint16) bool) UInter16
	EachReverse(func(int, uint16)) UInter16
	EachReverseBreak(func(int, uint16) bool) UInter16
	Fetch(int) uint16
	Get(int) (uint16, bool)
	Len() int
	Less(int, int) bool
	Map(func(int, uint16) uint16) UInter16
	Poll() uint16
	Pop() uint16
	Precatenate(UInter16) UInter16
	Prepend(...uint16) UInter16
	Push(...uint16) int
	Replace(int, uint16) bool
	Set() UInter16
	Sort() UInter16
	Swap(int, int)
	Unshift(...uint16) int
	Values() []uint16
}

// NewUInter16 returns a new UInter16 interface.
func NewUInter16(i ...uint16) UInter16 {
	return (&uinter16{&Slice{}}).Append(i...)
}

type uinter16 struct{ s *Slice }

func (u *uinter16) Append(i ...uint16) UInter16 {
	u.s.Append(uint16ToInterface(i...)...)
	return u
}

func (u *uinter16) Bounds(i int) bool {
	return u.s.Bounds(i)
}

func (u *uinter16) Concatenate(v UInter16) UInter16 {
	u.s.Concatenate(v.(*uinter16).s)
	return u
}

func (u *uinter16) Each(fn func(int, uint16)) UInter16 {
	u.s.Each(func(i int, v interface{}) {
		fn(i, (v.(uint16)))
	})
	return u
}

func (u *uinter16) EachBreak(fn func(int, uint16) bool) UInter16 {
	u.s.EachBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint16)))
	})
	return u
}

func (u *uinter16) EachReverse(fn func(int, uint16)) UInter16 {
	u.s.EachReverse(func(i int, v interface{}) {
		fn(i, (v.(uint16)))
	})
	return u
}

func (u *uinter16) EachReverseBreak(fn func(int, uint16) bool) UInter16 {
	u.s.EachReverseBreak(func(i int, v interface{}) bool {
		return fn(i, (v.(uint16)))
	})
	return u
}

func (u *uinter16) Fetch(i int) uint16 {
	var s, _ = u.Get(i)
	return s
}

func (u *uinter16) Get(i int) (uint16, bool) {
	var (
		ok bool
		s  uint16
	)
	ok = u.Bounds(i)
	if ok {
		s = (u.s.Fetch(i)).(uint16)
	}
	return s, ok
}

func (u *uinter16) Len() int {
	return (u.s.Len())
}

func (u *uinter16) Less(i int, j int) bool {
	return i < j
}

func (u *uinter16) Map(fn func(int, uint16) uint16) UInter16 {
	u.s.Map(func(i int, v interface{}) interface{} {
		return fn(i, (v.(uint16)))
	})
	return u
}

func (u *uinter16) Poll() uint16 {
	var (
		s uint16
		v = u.s.Poll()
	)
	if v != nil {
		s = (v.(uint16))
	}
	return s
}

func (u *uinter16) Pop() uint16 {
	var (
		s uint16
		v = u.s.Pop()
	)
	if v != nil {
		s = (v.(uint16))
	}
	return s
}

func (u *uinter16) Precatenate(v UInter16) UInter16 {
	u.s.Precatenate(v.(*uinter16).s)
	return u
}

func (u *uinter16) Prepend(i ...uint16) UInter16 {
	u.s.Prepend(uint16ToInterface(i...)...)
	return u
}

func (u *uinter16) Push(i ...uint16) int {
	return u.s.Push(uint16ToInterface(i...))
}

func (u *uinter16) Replace(i int, n uint16) bool {
	return (u.s.Replace(i, n))
}

func (u *uinter16) Set() UInter16 {
	u.s.Set()
	return u
}

func (u *uinter16) Sort() UInter16 {
	sort.Sort(u)
	return u
}

func (u *uinter16) Swap(i int, j int) {
	u.s.Swap(i, j)
}

func (u *uinter16) Unshift(i ...uint16) int {
	return (u.s.Unshift(uint16ToInterface(i...)))
}

func (u *uinter16) Values() []uint16 {
	var v = make([]uint16, u.Len())
	u.Each(func(i int, n uint16) {
		v[i] = n
	})
	return v
}

func uint16ToInterface(n ...uint16) []interface{} {
	var (
		i int
		v uint16
		x = make([]interface{}, (len(n)))
	)
	for i, v = range n {
		x[i] = v
	}
	return x
}
