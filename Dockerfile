FROM node:18.15.0 as jsbuild
WORKDIR /js/mock/
RUN git clone https://github.com/dhnikolas/jsmock.git . \
    && git checkout tags/v1.2 \
    && npm install --save react react-dom react-scripts \
    && REACT_APP_API_URL= npm run-script build

FROM golang:1.19 as build
WORKDIR /go/mock/
COPY . ./
RUN go mod vendor && GOOS=linux CGO_ENABLED=1 go build -o mockservice cmd/app/main.go

FROM golang:1.19
WORKDIR /go/mock/
RUN mkdir -p /var/mock/
COPY --from=build /go/mock/mockservice .
COPY --from=jsbuild /js/mock/build/static public/static
COPY --from=jsbuild /js/mock/build/index.html public/index.html

EXPOSE 8111

CMD ["./mockservice"]

