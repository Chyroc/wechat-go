test:
	go test ./...

test-v:
	go test -v ./...

codestyle:
	gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./ttt/*"`
	gofmt -d `find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./ttt/*"`
	goimports -w -local='github.com/Chyroc/qx-go' `find . -name '*.go' | grep -v vendor`
	golint -set_exit_status `go list ./... | grep -v /vendor/ | grep -v /ttt/*`

.PHONY: test
