# build stage
FROM golang:alpine AS builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git

RUN mkdir -p /go/src/github.com/prongbang/wiremock
WORKDIR /go/src/github.com/prongbang/wiremock
COPY . .

RUN go mod vendor

# With go ≥ 1.10
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/github.com/prongbang/wiremock

# final stage small image
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/github.com/prongbang/wiremock /go/bin/

USER appuser

ENTRYPOINT ["/go/bin/wiremock"]