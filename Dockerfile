FROM golang:1.17

WORKDIR /app
COPY . .
RUN make vendor \
    && make build

WORKDIR ./target/

CMD ["./goapi"]
