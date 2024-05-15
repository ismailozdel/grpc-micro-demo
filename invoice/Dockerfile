FROM golang:1.22.3-alpine3.19 as builder

WORKDIR /app

ENV CGO_=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build .

FROM scratch

COPY --from=builder /app /app

CMD ["./app/invoice"]