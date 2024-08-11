package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}
	readerData := bytes.NewReader([]byte("some data"))
	store := Store{
		StoreOpts: opts,
	}
	if err := store.writeStream("dataKey", readerData); err != nil {
		t.Error(err)
	}
}
