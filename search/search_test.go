// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package search

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewIndex(t *testing.T) {
	files := []string{
		"testdata/good", "testdata/headoverfeet", "testdata/ironic",
		"testdata/kingofpain", "testdata/pressure", "testdata/thankyou",
	}
	store, err := NewIndex(files...)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"testdata/good", "testdata/ironic", "testdata/thankyou"}
	if !reflect.DeepEqual(store.index["good"], want) {
		t.Errorf("NewIndex: Want %v. Got %v.", want, store.index["good"])
	}
	// Ignore marks, periods, dashes, etc.
	want = want[1:2]
	if !reflect.DeepEqual(store.index["ironic"], want) {
		t.Errorf("NewIndex: Want %v. Got %v.", want, store.index["ironic"])
	}
}

func TestNewIndexUnknownFile(t *testing.T) {
	_, err := NewIndex("testdata/good", "testdata/bad")
	if err == nil {
		t.Errorf("NewIndex. Want non-nil error, got <nil>.")
	}
}

func TestSearch(t *testing.T) {
	files := []string{
		"testdata/good", "testdata/headoverfeet", "testdata/ironic",
		"testdata/kingofpain", "testdata/pressure", "testdata/thankyou",
	}
	st, err := NewIndex(files...)
	if err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		input string
		want  []string
	}{
		{"good", []string{"testdata/good", "testdata/ironic", "testdata/thankyou"}},
		{"ironic", []string{"testdata/ironic"}},
		{"ironic | gained", []string{"testdata/good", "testdata/ironic"}},
		{"good & down", []string{"testdata/good", "testdata/ironic", "testdata/thankyou"}},
		{"good & down & ironic", []string{"testdata/ironic"}},
		{"secret", nil},
		{"secret | unsecret", nil},
		{"secret & ironic", nil},
		{"", nil},
	}
	for _, tt := range tests {
		got, err := st.Search(tt.input)
		sort.Strings(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Search(%q): Want %v. Got %v.", tt.input, tt.want, got)
		}
		if tt.want == nil && err == nil {
			t.Errorf("Expected non-nil error, got <nil>.")
		}
	}
}
