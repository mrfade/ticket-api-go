FROM golang:alpine AS builder
RUN apk update && apk add --no-cache 'git=~2'

ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/ticket-api-go/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server .

FROM alpine:3

WORKDIR /
COPY --from=builder /app/server /app/server

ENV PORT 5005
ENV GIN_MODE release
EXPOSE 5005

WORKDIR /app
ENTRYPOINT ["/app/server"]