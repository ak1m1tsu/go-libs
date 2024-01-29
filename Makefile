.PHONY: test
test:
	go test -v -race -coverprofile=cover.out ./... \
    	&& go tool cover -html=cover.out
