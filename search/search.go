// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package search

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

type Store struct {
	index map[string][]string
}

func NewStore(files ...string) (*Store, error) {
	s := Store{
		index: make(map[string][]string),
	}
	sort.Strings(files)
	for _, name := range files {
		f, err := os.Open(name)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			t := scanner.Text()
			t = strings.Trim(t, `.?!-,;:"`)
			l := s.index[t]
			if n := sort.SearchStrings(l, name); n == len(l) {
				l = append(l, name)
				s.index[t] = l
			}
		}
	}
	return &s, nil
}
