FROM golang:1.16-alpine AS BUILDER
WORKDIR /app

ADD src /app/src
COPY main.go ./
COPY go.mod ./
COPY go.sum ./

RUN go get

RUN go build -o /app/main main.go

FROM alpine
WORKDIR /app
COPY --from=BUILDER /app/main .
ENV PORT=80
CMD [ "/app/main" ]