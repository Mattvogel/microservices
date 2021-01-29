FROM golang:latest as builder

WORKDIR /go/src/devices
COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 go build -o /bin/devices

FROM scratch as runtime

COPY --chmod=111 --from=builder /bin/devices /bin/devices

CMD ["/bin/devices"]


