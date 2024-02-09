FROM golang:1.22.0-alpine

RUN go install github.com/cosmtrek/air@latest

COPY . ./app
WORKDIR /app

CMD [ "air", "-c", ".air.toml" ]