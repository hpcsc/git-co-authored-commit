package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	authors, readErr := ReadCoAuthors()
	if readErr != nil {
		log.Println(fmt.Errorf("error while reading .git-co-authors file: %v", readErr.Error()))
		os.Exit(1)
	}

	selected, selectErr := SelectCoAuthor(authors)
	if selectErr != nil {
		log.Println(fmt.Errorf("error while selecting co-author from list: %v", selectErr.Error()))
		os.Exit(1)
	}

	if selected == "" {
		// cancelled by user during selection
		os.Exit(0)
	}

	fmt.Println(selected)
}
