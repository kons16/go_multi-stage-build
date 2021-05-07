#######
# Build container
#######

FROM golang:1.14-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/hello_app
COPY ./services/hello_app .

RUN go get -u github.com/pilu/fresh && go build -o app

#######
# Exec container
#######

FROM scratch

EXPOSE 8000

COPY --from=builder /go/src/hello_app/app /app

CMD ["./app"]
