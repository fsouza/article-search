// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/fsouza/article-search/search"
	"log"
	"os"
	"strings"
)

func main() {
	var query string
	store, err := search.NewIndex(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Print("Search for: ")
		fmt.Scanf("%s", &query)
		articles, err := store.Search(query)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Matching articles: %s.\n", strings.Join(articles, ", "))
	}
}
