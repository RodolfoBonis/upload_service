FROM golang:1.19.2-alpine as build-env

ENV CGO_ENABLE 0

WORKDIR /go/src/github.com/RodolfoBonis/upload_service/
ADD . /go/src/github.com/RodolfoBonis/upload_service/

RUN go build -o upload_service -v /go/src/github.com/RodolfoBonis/upload_service/

COPY . ./

FROM alpine:3.15

WORKDIR /go/src/github.com/RodolfoBonis/upload_service/

COPY --from=build-env /go/src/github.com/RodolfoBonis/upload_service/upload_service /

CMD ["/upload_service"]

EXPOSE 8005