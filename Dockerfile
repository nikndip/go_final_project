FROM golang:1.27-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/todo .

FROM alpine:3.24

WORKDIR /app

COPY --from=builder /out/todo ./todo
COPY web ./web

RUN mkdir -p /data

ENV TODO_PORT=7540
ENV TODO_DBFILE=/data/scheduler.db

EXPOSE 7540

CMD ["./todo"]