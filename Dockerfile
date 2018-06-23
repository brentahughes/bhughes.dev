FROM golang:latest
COPY . /go/src/gitlab.com/bah2830/brentahughes.com
WORKDIR /go/src/gitlab.com/bah2830/brentahughes.com
RUN go build -o /app/personal-website


FROM scratch:latest
COPY --from=0 /app /app
COPY config.yaml /app/config.yaml
COPY templates /app/templates
EXPOSE 8080
CMD ["/app/personal-website"]
