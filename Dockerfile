FROM arm64v8/golang:1.23.0-alpine3.20

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

# VOLUME . /app

# RUN ls .

RUN go mod tidy

RUN go build -o binary

EXPOSE 8080

ENTRYPOINT ["/app/binary"]
