FROM golang:1.16.2-buster
WORKDIR /app
COPY . .
RUN go mod download
RUN go build
RUN mkdir /app/logs
EXPOSE 5000
CMD ["./hr-dms-api"]


