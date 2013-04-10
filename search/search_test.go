// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package search

import (
	"reflect"
	"testing"
)

func TestNewStore(t *testing.T) {
	files := []string{
		"testdata/good", "testdata/headoverfeet", "testdata/ironic",
		"testdata/kingofpain", "testdata/pressure", "testdata/thankyou",
	}
	store, err := NewStore(files...)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"testdata/good", "testdata/ironic", "testdata/thankyou"}
	if !reflect.DeepEqual(store.index["good"], want) {
		t.Errorf("NewStore: Want %v. Got %v.", want, store.index["good"])
	}
	// Ignore marks, periods, dashes, etc.
	want = want[1:2]
	if !reflect.DeepEqual(store.index["ironic"], want) {
		t.Errorf("NewStore: Want %v. Got %v.", want, store.index["ironic"])
	}
}

func TestNewStoreUnknownFile(t *testing.T) {
	_, err := NewStore("testdata/good", "testdata/bad")
	if err == nil {
		t.Errorf("NewStore. Want non-nil error, got <nil>.")
	}
}
