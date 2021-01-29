FROM golang:latest as builder

WORKDIR /go/src/auth
COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 go build -o /bin/auth

FROM scratch as runtime

COPY --chmod=111 --from=builder /bin/auth /bin/auth

CMD ["/bin/auth"]


