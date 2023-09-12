FROM golang:1.19.2-alpine as build-env

RUN apk add --no-cache git ca-certificates


ARG GITHUB_TOKEN
ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux TOKEN=$GITHUB_TOKEN

RUN git config --global url."https://${TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

WORKDIR /go/src/github.com/RodolfoBonis/upload_service/
ADD . /go/src/github.com/RodolfoBonis/upload_service/

RUN go env -w GOPRIVATE=github.com/RodolfoBonis/go_key_guardian

RUN go build -o upload_service -v /go/src/github.com/RodolfoBonis/upload_service/

COPY . ./

FROM alpine:3.15

WORKDIR /go/src/github.com/RodolfoBonis/upload_service/

COPY --from=build-env /go/src/github.com/RodolfoBonis/upload_service/upload_service /

CMD ["/upload_service"]

EXPOSE 8005