FROM golang:1.11
COPY . /go/src/github.com/bah2830/brentahughes.com
WORKDIR /go/src/github.com/bah2830/brentahughes.com
RUN go build -o /app/personal-website
EXPOSE 8080
CMD ["/app/personal-website"]
