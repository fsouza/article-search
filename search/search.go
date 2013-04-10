// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package search

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strings"
)

var noArticleErr = errors.New("No articles found.")

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

func (s *Store) search(word string) ([]string, error) {
	r, ok := s.index[word]
	if !ok {
		return nil, noArticleErr
	}
	return r, nil
}

func (s *Store) or(queries ...string) ([]string, error) {
	var result []string
	for _, q := range queries {
		q = strings.TrimSpace(q)
		r, _ := s.search(q)
		result = append(result, r...)
	}
	if result == nil {
		return nil, noArticleErr
	}
	return result, nil
}

func (s *Store) and(queries ...string) ([]string, error) {
	var articles []string
	for _, q := range queries {
		q = strings.TrimSpace(q)
		r, err := s.search(q)
		if err != nil {
			return nil, err
		}
		if articles == nil {
			articles = append(articles, r...)
		} else {
			cut := len(articles)
			for i := 0; i < cut; {
				a := articles[i]
				if n := sort.SearchStrings(r, a); n >= len(r) || r[n] != a {
					articles[i], articles[cut-1] = articles[cut-1], articles[i]
					cut--
				} else {
					i++
				}
			}
			articles = articles[:cut]
		}
	}
	return articles, nil
}

func (s *Store) Search(query string) ([]string, error) {
	parts := strings.Split(query, "&")
	if len(parts) > 1 {
		return s.and(parts...)
	}
	parts = strings.Split(query, "|")
	return s.or(parts...)
}
