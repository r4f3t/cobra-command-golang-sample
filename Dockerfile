FROM golang:1.17

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

EXPOSE 5558

# docker run -p 5556:5556  go-sample   go run main.go api -c ./config/config.qa.json -p 5556
