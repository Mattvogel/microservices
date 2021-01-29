FROM golang:latest as builder

WORKDIR /go/src/conditions
COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 go build -o /bin/conditions

FROM scratch as runtime

COPY --chmod=111 --from=builder /bin/conditions /bin/conditions

CMD ["/bin/conditions"]


