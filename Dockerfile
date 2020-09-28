FROM zgldh/docker-golang-builder:1.15.2-alpine3.12 AS build

WORKDIR /go/src
COPY app ./app
COPY main.go .
COPY go.mod .

ENV CGO_ENABLED=0
ENV GO111MODULE=on
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
RUN go get -v ./...

RUN go build -a -o main -ldflags '-extldflags "-static"' .

FROM scratch AS runtime
COPY --from=build /go/src/main ./
ENTRYPOINT ["./main"]
