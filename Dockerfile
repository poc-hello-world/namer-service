FROM golang:1.22.4 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

ADD . .

RUN go build \
  -ldflags "-X main.BuildID=${VERSION}" \
  -o namer

WORKDIR /dist
RUN cp /build/namer ./namer

FROM ubuntu:focal-20240530 AS app
COPY --from=builder /dist/namer /namer
ENTRYPOINT ["/namer"]
