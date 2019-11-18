# golang-tools

Demo how to install binary dependencies for Go with modules without abusing your code dependencies.

## How it works

* All binary dependencies described in `_tools/tools.go`. Symbol `_` hides code in `_tools` from go build tools.
* Directory `_tools` initialized as separate go module (in this case `github.com/nordicdyno/golang-tools-external`). This allows track binaries dependencies like it was proposed by «`tools.go` solution» in https://github.com/golang/go/issues/33326, but without a build tag.
* Binaries could be installed without any extra tools (except shell), just by running `_tools/install.sh`.
  * This script just runs `ls-imports` utility in the context of tools module and sends its output to `go install` command.
  * The tool `ls-imports` is carried inside `_tools/ls-imports` directory and all that it does just prints all imports from provided files.

## Why it's better than just putting tools.go inside your module context.

* Your project's binary tools and their dependencies separated from your code.
* Dependencies versions of these extra tools don't abuse your dependencies list and don't interfere in an unpredictable way with dependencies versions of your codebase.

## Why it's not an ideal solution

1. Dependencies of binary dependencies still interfere with each other.
    * It could be solved by separating tools into multiple modules.
2. You need a shell to install tools.
    * It's easily solvable with docker or a little bit more complicated `ls-imports` tool.

## How to use

build binary:

    go build

install tools (linter):

    make tools-install

run linter:

    make lint
