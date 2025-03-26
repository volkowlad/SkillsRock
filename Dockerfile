FROM golang:1.24
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o list-app ./cmd/main.go

CMD ["./list-app"]