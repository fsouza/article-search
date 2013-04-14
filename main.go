// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"github.com/fsouza/article-search/search"
	"log"
	"os"
	"strings"
)

func main() {
	store, err := search.NewIndex(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	fmt.Print("> ")
	for scanner.Scan() {
		articles, err := store.Search(scanner.Text())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s.\n", strings.Join(articles, ", "))
		}
		fmt.Print("> ")
	}
}
