FROM golang:alpine
COPY . .

RUN cd ./src && \
    go mod init main && \
    go get ./... && \
    go build main.go


FROM alpine
COPY --from=0 /go/src/main .
CMD ["./main"]
