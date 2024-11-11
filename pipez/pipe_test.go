package pipez

import (
	"reflect"
	"testing"
)

func TestPipe_Concat(t *testing.T) {
	var res = Of([]int{1, 2, 3}).Concat([]int{4, 5, 6}).Slice()
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}
}

func TestPipe_Concat_withoutWrap(t *testing.T) {
	var firstSlice Pipe[int] = []int{1, 2, 3}
	var secondSlice Pipe[int] = []int{4, 5, 6}
	var res []int = firstSlice.Concat(secondSlice)
	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Fail()
	}
}
