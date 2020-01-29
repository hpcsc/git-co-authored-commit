package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var gitCoauthorFileLocations = []string {
	".git/.git-co-authors",
	"~/.git-co-authors",
}

func ReadCoAuthors() ([]string, error) {
	path, findErr := findExistingCoauthorFile(gitCoauthorFileLocations)
	if findErr != nil {
		return nil, findErr
	}

	file, openErr := os.Open(path)
	if openErr != nil {
		return nil, openErr
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, scanner.Text())
		}
	}

	return lines, scanner.Err()
}

func findExistingCoauthorFile(fileLocations []string) (string, error) {
	for _, location := range fileLocations {
		expandedLocation := expandHomeDirectory(location)
		if fileExists(expandedLocation) {
			return expandedLocation, nil
		}
	}

	return "", errors.New(fmt.Sprintf("cannot find any .git-co-authors file, checked locations: %s", strings.Join(gitCoauthorFileLocations, ", ")))
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		// file may or may not exist. See err for details.
		return true
	}
}

func expandHomeDirectory(filePath string) string {
	usr, _ := user.Current()
	homeDirectory := usr.HomeDir
	if strings.HasPrefix(filePath, "~/") {
		return filepath.Join(homeDirectory, filePath[2:])
	}

	return filePath
}
