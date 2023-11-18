FROM golang:1.21.4 AS builder
RUN apt update
RUN apt install -y build-essential

WORKDIR /app
COPY . .
RUN make

FROM builder as lint
RUN make golangci

FROM builder as test
RUN make unit-tests


FROM ubuntu:22.04

RUN apt update && apt-get install -y --no-install-recommends ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/* .

EXPOSE 8080

CMD /app/server -config "/app"
