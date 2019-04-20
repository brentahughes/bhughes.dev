FROM golang:1.12
COPY . /goapp/
WORKDIR /goapp
RUN go mod tidy && go mod verify && go mod download && go build -o /app/personal-website
EXPOSE 8080
CMD ["/app/personal-website"]
