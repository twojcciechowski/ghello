
FROM golang:alpine AS builder
RUN adduser -D -g '' ghello
RUN mkdir -p /opt/ghello
RUN grep ghello /etc/passwd > /passwd1
WORKDIR /opt/ghello
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /ghello

FROM scratch
COPY --from=builder /passwd1 /etc/passwd
COPY --from=builder /ghello /ghello
USER ghello

ENTRYPOINT ["/ghello"]


