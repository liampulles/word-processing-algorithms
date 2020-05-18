# Keep test at the top so that it is default when `make` is called.
# This is used by Travis CI.
coverage.txt:
	mkdir -p /tmp/word-processing-algorithms
	go test -race -coverprofile=/tmp/word-processing-algorithms/pkg_coverage.txt -covermode=atomic -coverpkg=./pkg/... ./test/pkg/...
	cat /tmp/word-processing-algorithms/*_coverage.txt > coverage.txt
view-cover: coverage.txt
	go tool cover -html=coverage.txt
test: build
	go test ./test/...
build:
	go build ./...
install: build
	go install ./...
inspect: build
	golint ./...
pre-commit: clean coverage.txt inspect
	go mod tidy
clean:
	rm -f ${GOPATH}/bin/word-*
	rm -f coverage.txt
