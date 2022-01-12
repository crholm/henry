package slicez

import (
	"constraints"
	"errors"
	"github.com/modfin/go18exp/compare"
	"github.com/modfin/go18exp/slicez/sort"
	"math/rand"
	"time"
)

func Equal[A comparable](s1, s2 []A) bool {
	return EqualFunc(s1, s2, compare.Equal[A])
}
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

func Index[E comparable](s []E, needle E) int {
	return IndexFunc(s, func(e E) bool {
		return needle == e
	})
}

func IndexFunc[E any](s []E, f func(E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

func LastIndex[E comparable](s []E, needle E) int {
	return LastIndexFunc(s, func(e E) bool {
		return e == needle
	})
}
func LastIndexFunc[E any](s []E, f func(E) bool) int {
	n := len(s)

	for i := 0; i < n; i++ {
		if f(s[n-i-1]) {
			return n - i - 1
		}
	}
	return -1
}

func Cut[E comparable](s []E, needle E) (left, right []E, found bool) {
	return CutFunc(s, func(e E) bool {
		return e == needle
	})
}

func CutFunc[E any](s []E, on func(E) bool) (left, right []E, found bool) {
	i := IndexFunc(s, on)
	if i == -1 {
		return s, nil, false
	}
	return s[:i], s[i+1:], true
}

func Find[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := IndexFunc(s, equal)
	if i == -1 {
		return e, false
	}
	return s[i], true
}

func FindLast[E any](s []E, equal func(E) bool) (e E, found bool) {
	i := LastIndexFunc(s, equal)
	if i == -1 {
		return e, false
	}
	return s[i], true
}

func Join[E any](slices [][]E, sep []E) []E {
	if len(slices) == 0 {
		return []E{}
	}
	if len(slices) == 1 {
		return append([]E(nil), slices[0]...)
	}
	n := len(sep) * (len(slices) - 1)
	for _, v := range slices {
		n += len(v)
	}

	b := make([]E, n)
	bp := copy(b, slices[0])
	for _, v := range slices[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], v)
	}
	return b
}

func Contains[E comparable](s []E, v E) bool {
	return Index(s, v) >= 0
}
func ContainsFunc[E any](s []E, f func(e E) bool) bool {
	return IndexFunc(s, f) >= 0
}

func Clone[E any](s []E) []E {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}
	return append([]E{}, s...)
}
func Compare[E constraints.Ordered](s1, s2 []E) int {
	return CompareFunc(s1, s2, compare.Compare[E])
}
func CompareFunc[E1, E2 any](s1 []E1, s2 []E2, cmp func(E1, E2) int) int {
	s2len := len(s2)
	for i, v1 := range s1 {
		if i >= s2len {
			return +1
		}
		v2 := s2[i]
		if c := cmp(v1, v2); c != 0 {
			return c
		}
	}
	if len(s1) < s2len {
		return -1
	}
	return 0
}

func Concat[A any](slices ...[]A) []A {
	var ret []A
	for _, slice := range slices {
		ret = append(ret, slice...)
	}
	return ret
}

func Reverse[A any](slice []A) []A {
	l := len(slice)
	res := make([]A, l)
	for i, val := range slice {
		res[l-i-1] = val
	}
	return res
}

func Head[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[0], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

func Last[A any](slice []A) (A, error) {
	if len(slice) > 0 {
		return slice[len(slice)-1], nil
	}
	var zero A
	return zero, errors.New("slice does not have any elements")
}

func Nth[A any](slice []A, i int) A {
	var zero A
	n := len(slice)
	if n == 0 {
		return zero
	}
	if n == 1 {
		return slice[0]
	}

	i = i % n

	if i < 0 {
		i = len(slice) + i
	}
	return slice[i]
}

func Tail[A any](slice []A) []A {
	return Drop(slice, 1)
}

func Each[A any](slice []A, apply func(a A)) {
	for _, a := range slice {
		apply(a)
	}
}

func TakeWhile[A any](slice []A, take func(a A) bool) []A {
	var res []A
	for _, val := range slice {
		if !take(val) {
			break
		}
		res = append(res, val)
	}
	return res
}
func TakeRightWhile[A any](slice []A, take func(a A) bool) []A {
	var l = len(slice)
	var res []A
	for i := range slice {
		i = l - i - 1
		val := slice[i]
		if !take(val) {
			break
		}
		res = append([]A{val}, res...)
	}
	return res
}

func Take[A any](slice []A, i int) []A {
	var j int
	return TakeWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}
func TakeRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return TakeRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

func DropWhile[A any](slice []A, drop func(a A) bool) []A {
	if len(slice) == 0 {
		return nil
	}

	var index int = -1
	for i, val := range slice {
		if !drop(val) {
			break
		}
		index = i
	}

	var a []A

	if index == -1 {
		a = make([]A, len(slice))
		copy(a, slice)
		return a
	}

	if index+1 < len(slice) {
		a = make([]A, len(slice)-index-1)
		copy(a, slice[index+1:])
		return a
	}

	return a
}

func DropRightWhile[A any](slice []A, drop func(a A) bool) []A {

	if len(slice) == 0 {
		return nil
	}

	var index int = -1
	var l = len(slice)
	for i := range slice {
		i = l - i - 1
		val := slice[i]
		if !drop(val) {
			break
		}
		index = i
	}
	var a []A
	if index == -1 {
		a = make([]A, len(slice))
		copy(a, slice)
		return a
	}

	if 0 < index && index < len(slice) {
		a = make([]A, index)
		copy(a, slice[:index])
		return a
	}
	return a

}

func Drop[A any](slice []A, i int) []A {
	var j int
	return DropWhile(slice, func(_ A) bool {
		res := j < i
		j += 1
		return res
	})
}
func DropRight[A any](slice []A, i int) []A {
	i = len(slice) - i - 1
	j := len(slice) - 1
	return DropRightWhile(slice, func(_ A) bool {
		res := j > i
		j -= 1
		return res
	})
}

func Filter[A any](slice []A, include func(a A) bool) []A {
	var res []A
	for _, val := range slice {
		if include(val) {
			res = append(res, val)
		}
	}
	return res
}

func Reject[A any](slice []A, exclude func(a A) bool) []A {
	return Filter(slice, func(a A) bool {
		return !exclude(a)
	})
}

func Every[A comparable](slice []A, needle A) bool {
	return EveryFunc(slice, compare.EqualOf[A](needle))

}
func EveryFunc[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if !predicate(val) {
			return false
		}
	}
	return true
}

func Some[A comparable](slice []A, needle A) bool {
	return SomeFunc(slice, compare.EqualOf[A](needle))
}

func SomeFunc[A any](slice []A, predicate func(A) bool) bool {
	for _, val := range slice {
		if predicate(val) {
			return true
		}
	}
	return false
}

func None[A comparable](slice []A, needle A) bool {
	return !SomeFunc(slice, compare.EqualOf[A](needle))
}

func NoneFunc[A any](slice []A, predicate func(A) bool) bool {
	return !SomeFunc(slice, predicate)
}

func Partition[A any](slice []A, predicate func(a A) bool) (satisfied, notSatisfied []A) {
	for _, a := range slice {
		if predicate(a) {
			satisfied = append(satisfied, a)
			continue
		}
		notSatisfied = append(notSatisfied, a)
	}
	return satisfied, notSatisfied
}

func Shuffle[A any](slice []A) []A {
	var ret = append([]A{}, slice...)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return ret
}

func Sample[A any](slice []A, n int) []A {
	var ret []A

	if n > len(slice) {
		n = len(slice)
	}

	if n > len(slice)/3 { // sqare root?
		ret = Shuffle(slice)
		return ret[:n]
	}

	idxs := map[int]struct{}{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		var idx int
		for {
			idx = rand.Intn(len(slice))
			_, found := idxs[idx]
			if found {
				continue
			}
			idxs[idx] = struct{}{}
			break
		}

		ret = append(ret, slice[idx])
	}
	return ret
}

func Sort[A constraints.Ordered](slice []A) []A {
	return SortFunc(slice, compare.Less[A])
}
func SortFunc[A any](slice []A, less func(a, b A) bool) []A {
	var res = append([]A{}, slice...)
	sort.Slice(res, less)
	return res
}

func Search[A any](slice []A, f func(e A) bool) (index int, e A) {
	return sort.Search(slice, f)
}

func Compact[A comparable](slice []A) []A {
	return CompactFunc(slice, compare.Equal[A])
}

func CompactFunc[A any](slice []A, equal func(a, b A) bool) []A {
	if len(slice) == 0 {
		return slice
	}
	head := slice[0]
	last := head
	tail := Fold(slice[1:], func(accumulator []A, current A) []A {
		if equal(last, current) {
			return accumulator
		}
		last = current
		return append(accumulator, current)
	}, []A{})
	return append([]A{head}, tail...)
}

func Max[E constraints.Ordered](slice ...E) E {
	var zero E
	if slice == nil || len(slice) == 0 {
		return zero
	}
	cur := slice[0]
	for _, c := range slice {
		if cur < c {
			cur = c
		}
	}
	return cur
}

func Min[E constraints.Ordered](slice ...E) E {
	var zero E
	if slice == nil || len(slice) == 0 {
		return zero
	}
	cur := slice[0]
	for _, c := range slice {
		if cur > c {
			cur = c
		}
	}
	return cur
}

func Flatten[A any](slice [][]A) []A {
	var res []A
	for _, val := range slice {
		res = append(res, val...)
	}
	return res
}

func Map[A any, B any](slice []A, f func(a A) B) []B {
	res := make([]B, 0, len(slice))
	for _, a := range slice {
		res = append(res, f(a))
	}
	return res
}

func FlatMap[A any, B any](slice []A, f func(a A) []B) []B {
	return Flatten(Map(slice, f))
}

func Fold[I any, A any](slice []I, combined func(accumulator A, val I) A, init A) A {
	for _, val := range slice {
		init = combined(init, val)
	}
	return init
}

func FoldRight[I any, A any](slice []I, combined func(accumulator A, val I) A, init A) A {
	l := len(slice)
	for i := range slice {
		i := l - i - 1
		init = combined(init, slice[i])
	}
	return init
}

func KeyBy[A any, B comparable](slice []A, key func(a A) B) map[B]A {

	m := make(map[B]A)

	for _, v := range slice {
		k := key(v)
		m[k] = v
	}
	return m
}

func GroupBy[A any, B comparable](slice []A, key func(a A) B) map[B][]A {

	m := make(map[B][]A)

	for _, v := range slice {
		k := key(v)
		m[k] = append(m[k], v)
	}
	return m
}

func Uniq[A comparable](slice []A) []A {
	return UniqBy(slice, compare.Identity[A])
}

func UniqBy[A any, B comparable](slice []A, by func(a A) B) []A {
	var res []A
	var set = map[B]struct{}{}
	for _, e := range slice {
		key := by(e)
		_, ok := set[key]
		if ok {
			continue
		}
		set[key] = struct{}{}
		res = append(res, e)
	}
	return res
}

func Union[A comparable](slices ...[]A) []A {
	return UnionBy(compare.Identity[A], slices...)
}

func UnionBy[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var res []A
	var set = map[B]struct{}{}
	for _, slice := range slices {
		for _, e := range slice {
			key := by(e)
			_, ok := set[key]
			if ok {
				continue
			}
			set[key] = struct{}{}
			res = append(res, e)
		}
	}
	return res
}

func Intersection[A comparable](slices ...[]A) []A {
	return IntersectionBy(compare.Identity[A], slices...)
}
func IntersectionBy[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var res = UniqBy(slices[0], by)
	for _, slice := range slices[1:] {
		var set = map[B]bool{}
		for _, e := range slice {
			set[by(e)] = true
		}
		res = Filter(res, func(a A) bool {
			return set[by(a)]
		})
	}
	return res
}

func Difference[A comparable](slices ...[]A) []A {
	return DifferenceBy(compare.Identity[A], slices...)
}
func DifferenceBy[A any, B comparable](by func(a A) B, slices ...[]A) []A {
	if len(slices) == 0 {
		return nil
	}
	var exclude = map[B]bool{}
	for _, v := range IntersectionBy(by, slices...) {
		exclude[by(v)] = true
	}

	var res []A
	for _, slice := range slices {
		for _, e := range slice {
			key := by(e)
			if exclude[key] {
				continue
			}
			exclude[key] = true
			res = append(res, e)
		}
	}
	return res
}

func Complement[A comparable](a, b []A) []A {
	return ComplementBy(compare.Identity[A], a, b)
}
func ComplementBy[A any, B comparable](by func(a A) B, a, b []A) []A {
	if len(a) == 0 {
		return b
	}

	var exclude = map[B]bool{}
	for _, e := range a {
		exclude[by(e)] = true
	}

	var res []A
	for _, e := range b {
		key := by(e)
		if exclude[key] {
			continue
		}
		exclude[key] = true
		res = append(res, e)
	}

	return res
}

func Zip[A any, B any, C any](aSlice []A, bSlice []B, zipper func(a A, b B) C) []C {
	var i = len(aSlice)
	var j = len(bSlice)
	if j < i {
		i = j
	}
	var cSlice []C
	for k, a := range aSlice {
		if k == j {
			break
		}
		b := bSlice[k]
		cSlice = append(cSlice, zipper(a, b))
	}
	return cSlice
}

func Unzip[A any, B any, C any](cSlice []C, unzipper func(c C) (a A, b B)) ([]A, []B) {
	var aSlice []A
	var bSlice []B
	for _, c := range cSlice {
		a, b := unzipper(c)
		aSlice = append(aSlice, a)
		bSlice = append(bSlice, b)
	}
	return aSlice, bSlice

}
