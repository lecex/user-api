FROM bigrocs/golang-gcc:1.13 as builder

WORKDIR /go/src/github.com/lecex/user-api
COPY . .

ENV GO111MODULE=off CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -o bin/service

FROM bigrocs/alpine:ca-data

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=builder /go/src/github.com/lecex/user-api/bin/service /usr/local/bin/
CMD ["service"]