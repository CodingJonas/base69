# Golang Base69
This is a implementation of the `base69` encoding in Go, using only the standard library. This makes it easy to build and update.

## Usage
Run `base69 --help` to see the usage.

Here is an example on how to use this:
```sh
$ base69 'sixtynine!'
5AVBvAHAjAgBXBkB3AZAkAQAAAAAAA4=
$ base69 -d '5AVBvAHAjAgBXBkB3AZAkAQAAAAAAA4='
sixtynine!
```

## Installation
Tested on Ubuntu 24.04 using go1.22.

You can build this project using the Makefile. The executable will be at `./exec/base69`
```sh
make build
```
