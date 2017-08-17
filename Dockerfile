FROM golang:latest

COPY . /go/src/github.com/bah2830/brentahughes.com

WORKDIR /go/src/github.com/bah2830/brentahughes.com

RUN go get
RUN go build -o /app/personal-website

EXPOSE 80

CMD ["/app/personal-website"]
