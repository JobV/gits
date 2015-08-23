# Gits

Gits is a simple Go wrapper for Git commands on the command line.

All it does is run git commands wherever you execute your
go code.

## Install

```
go get github.com/JobV/gits
```

## Use

Clone a git repository:

```
err := gits.Clone(CLONE_URL string, TARGET_PATH string)

# For example:
err := gits.Clone("git@github.com:JobV/gits.git", "~/gits")
```
