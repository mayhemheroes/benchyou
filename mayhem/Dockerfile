# FROM fuzzers/go-fuzz:1.2.0
FROM golang:1.23 AS builder

COPY . /benchyou
WORKDIR /benchyou/mayhem

# Install go-fuzz tools using a local module to pin golang.org/x/tools to a
# version compatible with Go 1.23 (go-fuzz@latest pulls x/tools@v0.44.0
# which requires Go 1.25.0, so we use go build in a local module instead).
RUN mkdir /tmp/gofuzz && cd /tmp/gofuzz && \
    go mod init gofuzz && \
    GOFLAGS=-mod=mod go get golang.org/x/tools@v0.21.0 && \
    GOFLAGS=-mod=mod go get github.com/dvyukov/go-fuzz@v0.0.0-20240924070022-e577bee5275c && \
    GOFLAGS=-mod=mod go get github.com/elazarl/go-bindata-assetfs && \
    GOFLAGS=-mod=mod go get github.com/stephens2424/writerset && \
    GOFLAGS=-mod=mod go build -o /go/bin/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz && \
    GOFLAGS=-mod=mod go build -o /go/bin/go-fuzz-build github.com/dvyukov/go-fuzz/go-fuzz-build

RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzz.a && \
    clang -fsanitize=fuzzer fuzz.a -o fuzz.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /benchyou/mayhem/fuzz.libfuzzer /
