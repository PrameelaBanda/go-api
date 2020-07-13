FROM golang:1.14-alpine3.11
COPY /bin/walletservice /walletservice
RUN chmod u+x /go-api
ENTRYPOINT ["/go-api"]