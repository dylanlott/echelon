package main

import (
	"testing"

	"github.com/matryer/is"
)

type MockItem struct {
	val string
}

func TestQueue(t *testing.T) {
	is := is.New(t)

	q := &Queue[*MockItem]{}

	_, err := q.Add(&MockItem{val: "foo"})
	is.NoErr(err)
	_, err = q.Add(&MockItem{val: "bar"})
	is.NoErr(err)
	_, err = q.Add(&MockItem{val: "buz"})
	is.NoErr(err)

	is.Equal(len(q.items), 3)

	ok, peeked := q.Peek()
	is.True(ok)
	is.Equal(peeked.val, "foo")

	got, err := q.Pop()
	is.NoErr(err)
	is.Equal(got.val, "foo")

	got, err = q.Pop()
	is.NoErr(err)
	is.Equal(got.val, "bar")

	ok, peeked = q.Peek()
	is.True(ok)
	is.Equal(peeked.val, "buz")
}
