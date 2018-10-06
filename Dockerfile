# Stage 1. Build the binary
FROM golang:1.11 as builder

ENV RELEASE "0.0.3"
RUN adduser -u 10001 myuser

ENV WORKSPACE /go/src/github.com/makpoc/go-k8s-workshop
RUN mkdir -p $WORKSPACE

ADD . $WORKSPACE

WORKDIR $WORKSPACE

RUN go get ./... && \
  CGO_ENABLED=0 go build \
    -ldflags "-s -w -X github.com/makpoc/go-k8s-workshop/internal/version.Version=${RELEASE}" \
    -o bin/go-k8s-workshop github.com/makpoc/go-k8s-workshop/cmd/go-k8s-workshop

# Stage 2. Run the binary
FROM scratch

ENV PORT 8888
ENV DIAG_PORT 9999

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

USER myuser

COPY --from=builder /go/src/github.com/makpoc/go-k8s-workshop/bin/go-k8s-workshop ./go-k8s-workshop

EXPOSE $PORT
EXPOSE $DIAG_PORT

CMD ["/go-k8s-workshop"]
