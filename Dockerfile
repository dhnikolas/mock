FROM golang:latest as build
WORKDIR /go/mock/
COPY . .
RUN go mod vendor && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o mockservice cmd/app/main.go

FROM alpine:latest
WORKDIR /go/mock/
COPY --from=build /go/mock/mockservice .

CMD ["./mockservice"]

