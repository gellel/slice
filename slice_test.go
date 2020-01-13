package slice_test

import (
	"testing"

	"github.com/gellel/slice"
)

var (
	b   slice.Byter
	f32 slice.Floater32
	f64 slice.Floater64
	i   slice.Inter
	i8  slice.Inter8
	i16 slice.Inter16
	i32 slice.Inter32
	i64 slice.Inter64
	s   *slice.Slice
	u   slice.UInter
	u8  slice.UInter8
	u16 slice.UInter16
	u32 slice.UInter32
	u64 slice.UInter64
	v   slice.Interfacer
)

func Test(t *testing.T) {
	s = &slice.Slice{}
}

func TestAppend(t *testing.T) {
	if ok := (*s.Append("a"))[0].(string) == "a"; !ok {
		t.Fatalf("(&slice.Slice.Append(interface{})) != (interface{}))")
	}
}

func TestBounds(t *testing.T) {
	if ok := s.Bounds(s.Len() - 1); !ok {
		t.Fatalf("(&slice.Slice.Bounds(int) bool) != true")
	}
}

func TestConcatenate(t *testing.T) {
	if ok := (*s.Concatenate(&slice.Slice{"b"}))[1].(string) == "b"; !ok {
		t.Fatalf("(&slice.Slice.Concatenate(interface{})) != (interface{}))")
	}
}

func TestEach(t *testing.T) {
	s.Each(func(i int, v interface{}) {
		if ok := (*s)[i] == v; !ok {
			t.Fatalf("(&slice.Slice.Each(int, interface{})) != (interface{})")
		}
	})
}

func TestEachBreak(t *testing.T) {
	var (
		n int
	)
	s.EachBreak(func(i int, _ interface{}) bool {
		n = i
		return false
	})
	if ok := n == 0; !ok {
		t.Fatalf("(&slice.Slice.EachBreak(int, interface{}) bool) != true")
	}
}

func TestEachReverse(t *testing.T) {
	var (
		n = s.Len() - 1
	)
	s.EachReverse(func(i int, _ interface{}) {
		if ok := i == n; !ok {
			t.Fatalf("(&slice.Slice.EachReverse(i int, interface{})) > s.Len() - i")
		}
		n = n - i
	})
}
