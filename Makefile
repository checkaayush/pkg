PHONY: lint test coverage

lint:
	golangci-lint run

test:
	# Stop at first failing test. Refer: https://github.com/golang/go/issues/33038.
	# -race requires cgo; hence CGO_ENABLED=1
	for s in $$(go list ./...); do if ! CGO_ENABLED=1 go test -race -failfast -v -p 1 $$s; then break; fi; done

coverage:
	CGO_ENABLED=1 go test -race -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
	printf "\nTotal Unit Test Coverage: " && go tool cover -func coverage.out | grep total | awk '{print $$3}'