FROM golang:latest

COPY . /go/src/gitlab.com/bah2830/personal-website

WORKDIR /go/src/gitlab.com/bah2830/personal-website

RUN go get
RUN go build -o /app/personal-website

EXPOSE 80

CMD ["/app/personal-website"]
