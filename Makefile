build:
	go build -o bin/git-co-authored-commit

run: build
	./bin/git-co-authored-commit $(message)
