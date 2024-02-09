FROM golang:1.22.0-alpine

COPY . /app
WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

CMD [ "/app/api" ]