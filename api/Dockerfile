# FROM golang:alpine as builder

# RUN apk update && apk add --no-cache git

# WORKDIR /app

# COPY go.mod go.sum ./


# RUN go mod download 

# COPY . .

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates

# WORKDIR /root/

# COPY --from=builder /app/main .
# COPY --from=builder /app/.env .  

FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

#stage 2

FROM alpine

RUN adduser -S -D -H  -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]



