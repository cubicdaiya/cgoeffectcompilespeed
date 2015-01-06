all: go/go cgo/cgo cgofast/cgofast

go/go:
	cd go && time go build

cgo/cgo:
	cd cgo && time go build

cgofast/cgofast:
	cd cgofast && time go build

clean:
	rm -rf go/go cgo/cgo cgofast/cgofast
