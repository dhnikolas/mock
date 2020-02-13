FROM node:12.2.0 as jsbuild
WORKDIR /js/mock/
RUN git clone https://github.com/dhnikolas/jsmock.git . \
    && git checkout tags/v1.0 \
    && npm install --save react react-dom react-scripts \
    && REACT_APP_API_URL= npm run-script build

FROM golang:latest as build
WORKDIR /go/mock/
RUN git clone https://github.com/dhnikolas/mock.git . \
    && git checkout tags/v1.2 \
    && go mod vendor && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o mockservice cmd/app/main.go

FROM alpine:latest
WORKDIR /go/mock/
COPY --from=build /go/mock/mockservice .
COPY --from=jsbuild /js/mock/build/static public/static
COPY --from=jsbuild /js/mock/build/index.html public/index.html

EXPOSE 8111

CMD ["./mockservice"]

