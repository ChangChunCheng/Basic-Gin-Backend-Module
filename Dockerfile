FROM golang:1.16.3

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN make build

EXPOSE 8080

ENTRYPOINT [ "./app" ]