FROM golang:1.4.2-wheezy

RUN mkdir -p /app
WORKDIR /app
COPY . /app/
ENV GOPATH /go/
RUN go get -d -v
EXPOSE 8080

CMD bash -c 'go run *.go'
