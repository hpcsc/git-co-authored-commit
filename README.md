# git-co-authored-commit

[![Demo](https://github.com/hpcsc/git-co-authored-commit/raw/master/git-co-authored-commit.gif)](https://github.com/hpcsc/git-co-authored-commit/raw/master/git-co-authored-commit.gif)

This is the Golang version of my shell script [fcm](https://github.com/hpcsc/dotfiles/blob/master/link/common/zsh/.functions/fzf-functions/fcm)

The shell script version uses [fzf](https://github.com/junegunn/fzf) to interactively select a co-author from a predefined list of co-authors and invoke `git commit` with that co-author as [co-authored-by](https://help.github.com/en/github/committing-changes-to-your-project/creating-a-commit-with-multiple-authors) trailer. This `co-authored-by` trailer is useful in teams practicing pair programming.

The shell script above only works in `zsh` and MacOS/Linux environment. This repository is an attempt to make a cross-platform version of that using Golang.

## Installation

- Download binary for your platform from [Github releases](https://github.com/hpcsc/git-co-authored-commit/releases)
- For MacOS/Linux, `chmod +x` and move file to a location in your `PATH` (optionally rename it to something shorter, e.g. `gcam`)

  For Windows, add `.exe` extension to the file, optionally rename it to something shorter, add file location to your `PATH` variable
- Create `.git-co-authors` file at your home directory (for Windows, it's under `C:/Users/your-user`) or in your `.git` folder in your repository, each line is one co-author, .e.g.

```
author-1 <author1@email.com>
author-2 <author2@email.com>
```

## Limitation

It doesn't support multiple select like `fzf`. I tried looking for a Golang TUI library that supports multi-select dropdown but haven't found any.
