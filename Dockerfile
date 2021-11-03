FROM golang:1.17.2-alpine3.14 AS GO_BUILD
COPY . /server
WORKDIR /server/server
RUN go build -o /go/bin/server/server

FROM alpine:3.10
WORKDIR /agendapro
COPY --from=GO_BUILD /go/bin/server/ ./
EXPOSE 8085
CMD ./server
