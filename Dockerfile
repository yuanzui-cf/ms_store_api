FROM rust:slim-bookworm AS builder

WORKDIR /app

COPY Cargo.toml Cargo.lock ./

COPY msstore ./msstore
COPY server ./server

RUN cargo build --release

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/target/release/msapi ./server

EXPOSE 9000

CMD ["./server"]
