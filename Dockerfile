FROM golang:1.23-alpine

RUN apk update && apk upgrade && apk add --no-cache ca-certificates openssl

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o home_rep_cloud .

EXPOSE ${API_PORT}

CMD ["./home_rep_cloud"]
