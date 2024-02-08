FROM golang:1.22.0-alpine

COPY . ./app

WORKDIR /app

RUN go mod download
RUN go build -o api .

CMD [ "/api" ]