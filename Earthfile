VERSION 0.6


ARG GO_VERSION=1.16
ARG GOLANGCI_VERSION=1.40.1

base-go:
    FROM golang:$GO_VERSION
    
    RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg/mod -- \
        go install github.com/axw/gocov/gocov@latest \
        && go install github.com/AlekSi/gocov-xml@latest \
        && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v${GOLANGCI_VERSION}

    ENV CGO_ENABLED=0
    ENV GOOS=linux
    ENV GOARCH=amd64

    WORKDIR /build

copy-code:
    FROM +base-go
    COPY --dir */. *go* .golangci.yml ./

deps:
    FROM +copy-code
    COPY go.mod go.sum .
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

lint:  ## Lints the go code with various linters
    FROM +deps
    ## This linter is controlled via a config file, and is super-strict
    RUN --mount=type=cache,target=/root/.cache/golangci-lint \
        --mount=type=cache,target=/root/.cache/go-build -- \
        golangci-lint run -v --print-resources-usage \
	--modules-download-mode=readonly --timeout=10m 

test:  ## Run tests and save coverage
    FROM +deps
    RUN --mount=type=cache,target=/root/.cache/go-build -- \
        go test -v -mod=readonly -coverprofile=cover.out ./... \
        && gocov convert cover.out | gocov-xml > coverage.xml
    SAVE ARTIFACT coverage.xml AS LOCAL coverage.xml


lintest: ## One step for lint and test
    BUILD +lint
    BUILD +test
