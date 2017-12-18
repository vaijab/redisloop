FROM golang:1.9

WORKDIR /go/src/github.com/vaijab/redisloop
COPY . ./
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go test -v github.com/vaijab/redisloop
RUN go vet github.com/vaijab/redisloop
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /build/redisloop github.com/vaijab/redisloop


FROM scratch
COPY --from=0 /build/redisloop /redisloop
USER 1000
ENTRYPOINT ["/redisloop"]
