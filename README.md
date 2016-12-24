yesman
====

yesman はgo言語で書かれたCLIツールです。

## Description
「スクリプトへ繰り返し同じ入力をしなくてはならない」時にパイプラインで入力を渡すような用途を考えています。例えば `yesman "passwd" | ./install.sh` のような使い方です。

## Requirement
go : version 1.6 以上

## Usage
`$ yesman [-count=<int>] <word>`

## Install
```
$ go get github.com/byakheee/yesman
$ cd $GOPATH/src/github.com/byakheee/yesman
$ make setup
$ make build
$ export PATH=$PATH:$GOPATH/src/github.com/byakheee/yesman/bin
$ yesman -version
version: v0.1.0, revision: <rev>
```

### make help
```
$ make help
build:             build binaries
build-debug:       build debug binaries
deps:              Install dependencies
fmt:               Formmat source codes
help:              Show help
lint:              Lint
setup:             Setup
test:              Run tests
update:            Update dependencies
```

## Licence
The MIT License (MIT)

Copyright (c) 2016 byakheee

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

## Author
[byakheee](https://github.com/byakheee)