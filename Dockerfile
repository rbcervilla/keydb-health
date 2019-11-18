FROM golang:1.12.10 as builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o keydb-health

FROM scratch
COPY --from=builder /build/keydb-health .
ENTRYPOINT ["/keydb-health"]