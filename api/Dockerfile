FROM golang:1.16

WORKDIR /golinks

COPY . .

RUN go mod tidy
RUN go build server.go

ENTRYPOINT [ "/golinks/server" ]

EXPOSE 1323