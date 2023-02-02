FROM golang:1.19-alpine
WORKDIR /go/src/app
COPY . .

EXPOSE 3003:3003
CMD go run .
