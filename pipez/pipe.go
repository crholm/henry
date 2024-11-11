package pipez

import (
	"github.com/modfin/henry/slicez"
)

func Of[A any](a []A) Pipe[A] {
	return a
}

type Pipe[A any] []A

func (p Pipe[A]) Slice() []A {
	return p
}

func (p Pipe[A]) Peek(apply func(a A)) Pipe[A] {
	slicez.ForEach(p, apply)
	return p
}

func (p Pipe[A]) Concat(slices ...[]A) Pipe[A] {
	return slicez.Concat(append([][]A{p}, slices...)...)
}

func (p Pipe[A]) Tail() Pipe[A] {
	return slicez.Tail(p)
}
func (p Pipe[A]) Head() (A, error) {
	return slicez.Head(p)
}
func (p Pipe[A]) Last() (A, error) {
	return slicez.Last(p)
}
func (p Pipe[A]) Initial() Pipe[A] {
	return slicez.Initial(p)
}

func (p Pipe[A]) Reverse() Pipe[A] {
	return slicez.Reverse(p)
}

func (p Pipe[A]) Nth(i int) A {
	return slicez.Nth(p, i)
}

func (p Pipe[A]) Take(i int) Pipe[A] {
	return slicez.Take(p, i)
}
func (p Pipe[A]) TakeRight(i int) Pipe[A] {
	return slicez.TakeRight(p, i)
}

func (p Pipe[A]) TakeWhile(take func(a A) bool) Pipe[A] {
	return slicez.TakeWhile(p, take)
}
func (p Pipe[A]) TakeRightWhile(take func(a A) bool) Pipe[A] {
	return slicez.TakeRightWhile(p, take)
}

func (p Pipe[A]) Drop(i int) Pipe[A] {
	return slicez.Drop(p, i)
}

func (p Pipe[A]) DropRight(i int) Pipe[A] {
	return slicez.DropRight(p, i)
}

func (p Pipe[A]) DropWhile(drop func(a A) bool) Pipe[A] {
	return slicez.DropWhile(p, drop)
}

func (p Pipe[A]) DropRightWhile(drop func(a A) bool) Pipe[A] {
	return slicez.DropRightWhile(p, drop)
}

func (p Pipe[A]) Filter(include func(a A) bool) Pipe[A] {
	return slicez.Filter(p, include)
}

func (p Pipe[A]) Reject(exclude func(a A) bool) Pipe[A] {
	return slicez.Reject(p, exclude)
}

func (p Pipe[A]) Map(f func(a A) A) Pipe[A] {
	return slicez.Map(p, f)
}

func (p Pipe[A]) Fold(combined func(accumulator A, val A) A, accumulator A) A {
	return slicez.Fold(p, combined, accumulator)
}

func (p Pipe[A]) FoldRight(combined func(accumulator A, val A) A, accumulator A) A {
	return slicez.Fold(p, combined, accumulator)
}

func (p Pipe[A]) Every(predicate func(a A) bool) bool {
	return slicez.EveryBy(p, predicate)
}
func (p Pipe[A]) Some(predicate func(a A) bool) bool {
	return slicez.SomeBy(p, predicate)
}
func (p Pipe[A]) None(predicate func(a A) bool) bool {
	return slicez.NoneBy(p, predicate)
}

func (p Pipe[A]) Partition(predicate func(a A) bool) (satisfied, notSatisfied []A) {
	return slicez.Partition(p, predicate)
}

func (p Pipe[A]) Sample(n int) Pipe[A] {
	return slicez.Sample(p, n)
}
func (p Pipe[A]) Shuffle() Pipe[A] {
	return slicez.Shuffle(p)
}
func (p Pipe[A]) SortFunc(less func(a, b A) bool) Pipe[A] {
	return slicez.SortBy(p, less)
}

func (p Pipe[A]) Compact(equal func(a, b A) bool) Pipe[A] {
	return slicez.CompactBy(p, equal)
}

func (p Pipe[A]) Count() int {
	return len(p)
}

func (p Pipe[A]) Zip(b []A, zipper func(a, b A) A) Pipe[A] {
	return slicez.Zip(p, b, zipper)
}
func (p Pipe[A]) Unzip(unzipper func(a A) (A, A)) ([]A, []A) {
	return slicez.Unzip(p, unzipper)
}

func (p Pipe[A]) Interleave(a ...[]A) Pipe[A] {
	return slicez.Interleave[A](append([][]A{p}, a...)...)
}
