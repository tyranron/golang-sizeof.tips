Golang sizeof tips
==================

_Moved from [gophergala/golang-sizeof.tips](https://github.com/gophergala/golang-sizeof.tips)_

**Web tool for interactive playing with Golang struct sizes.**

Try online version [here](http://golang-sizeof.tips/).

## Aim
Provide comfortable tool to see how fields in struct are aligned,
to compare different structs and as the result - to understand
and remember alignment rules.

## Installing
```bash
go get github.com/gophergala/golang-sizeof.tips
cd github.com/gophergala/golang-sizeof.tips
go build -o ./server
```
You may also install via simple `go get` by your own risk.


## Usage
```bash
./server -http=:7777 start
./server stop
./server restart
```

## Platform support
Tested on Linux and OS X x64 platforms, but should work properly and on other
*nix-like platforms.
Windows is not supported due to daemonization.

## License
[Apache License 2.0](LICENSE)
