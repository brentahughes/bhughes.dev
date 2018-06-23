FROM golang:latest
COPY . /go/src/gitlab.com/bah2830/brentahughes.com
WORKDIR /go/src/gitlab.com/bah2830/brentahughes.com
RUN go build -o /app/personal-website
EXPOSE 8080
CMD ["/app/personal-website"]
