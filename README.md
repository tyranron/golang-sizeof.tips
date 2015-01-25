Golang sizeof tips
------------------

**Web tool for interactive playing with Golang struct sizes.**

## Install
To install correct versions of dependencies
[Goop dependency manager](https://github.com/nitrous-io/goop) should be used.
```bash
go get github.com/gophergala/golang-sizeof.tips
cd github.com/gophergala/golang-sizeof.tips
goop install
goop go build -o ./server
```
You may also install via simple `go get` by your own risk.


## Starting, stoping, restarting
```bash
./server -http=:7777 start
./server stop
./server restart
```
