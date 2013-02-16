all:

test:
	GOPATH=$(realpath .) go test -test.v judy

