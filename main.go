package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

	message := getCommitMessageFromArgument()
	commitOutput, commitErr := commit(message, selected)
	if commitErr != nil {
		log.Println(fmt.Errorf("error while committing: %v", commitErr))
	}

	fmt.Println(string(commitOutput))
}

func getCommitMessageFromArgument() string {
	return strings.Join(os.Args[1:], " ")
}

func commit(message string, coauthor string) ([]byte, error) {
	// message variable can contains \n characters. Those characters are escaped by Golang so they will not be rendered as newline
	// in the final message. To render those \n characters, use strings.Replace() to replace escaped newline with unescaped
	// newline. Strings wrapped in backticks can be written over multiple lines so will be escaped by Golang
	messageWithCoauthor := fmt.Sprintf("%s\n\nco-authored-by: %s", strings.Replace(message, `\n`, "\n", -1), coauthor)

	command := exec.Command("git", "commit", "-m", messageWithCoauthor)
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	return command.Output()
}