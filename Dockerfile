FROM golang:1.22

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -a -o muzz-api

EXPOSE 3000

ENTRYPOINT [ "/app/muzz-api" ]
