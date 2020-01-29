# git-co-authored-commit

This is the Golang version of my shell script [fcm](https://github.com/hpcsc/dotfiles/blob/master/link/common/zsh/.functions/fzf-functions/fcm)

The shell script version uses [fzf](https://github.com/junegunn/fzf) to interactively select a co-author from a predefined list of co-authors and invoke `git commit` with that co-author as [co-authored-by](https://help.github.com/en/github/committing-changes-to-your-project/creating-a-commit-with-multiple-authors) trailer. This `co-authored-by` trailer is useful in teams practicing pair programming.

The shell script above only works in `zsh` and MacOS/Linux environment. This repository is an attempt to make a cross-platform version of that using Golang.
