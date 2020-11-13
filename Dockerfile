FROM golang:alpine
COPY main.go .
ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

RUN  go mod init \
    && go mod download\
    && go build main.go

FROM alpine
COPY --from=0 /go/main .
CMD ["./main"]