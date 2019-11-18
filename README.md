# golang-tools

Demo how to install binary dependencies for Go with modules without abusing your code dependencies.

## How it works

* binary dependencies described in `_tools/linter/tools.go`.
    * symbol `_` in top level directory name hides dependency code from go build tools.
    * second level directory `linter` exist because it's allows go tools work properly inside it (go tools don't like directory names with underscores)
* Directory `_tools/linter` initialized as separate go module (in this case `github.com/nordicdyno/golang-tools-linter`). This allows track binaries dependencies like it was proposed by «`tools.go` solution» in https://github.com/golang/go/issues/25922, but without mixing tools and project dependencies.
* Binaries could be installed without any extra tools (except shell), just by running `_tools/install.sh`.
  * This script just runs `ls-imports` utility in the context of tools module and sends its output to `go install` command.
  * The tool `ls-imports` is carried inside `_tools/ls-imports` directory and all that it does just prints all imports from provided files.

## Why it's better than just putting tools.go inside your module context.

* Your project's binary tools and their dependencies separated from your code.
* Versions of dependencies these extra tools don't collide with your dependencies versions and don't interfere with them in unpredictable ways.

## Why it's not an ideal solution

1. Dependencies of binary dependencies still interfere with each other, until we separate them in their own `dependency holder` modules
2. You need a shell to install tools. (It's easily solvable with docker or a little bit more complicated `ls-imports`/`deps-install` tool)

## How to use

build binary:

    go build

install tools (linter):

    make tools-install

run linter:

    make lint
