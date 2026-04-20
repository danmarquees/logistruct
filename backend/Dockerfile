FROM golang:1.22-bookworm AS builder

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

WORKDIR /app

# Copy Rust core
COPY core ./core/
# Pre-build rust core to cache it
RUN cd core && cargo build --release

# Copy Go backend
COPY backend ./backend/
WORKDIR /app/backend

# Initialize go modules if not already
RUN go mod tidy

# Build Go application
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o logistruct-backend ./cmd/server

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates libc6 && rm -rf /var/lib/apt/lists/*

WORKDIR /app
# Copy the rust library globally so it can be dynamically linked
COPY --from=builder /app/core/target/release/liblogistruct_core.so /usr/lib/liblogistruct_core.so
COPY --from=builder /app/backend/logistruct-backend .

EXPOSE 8080

CMD ["./logistruct-backend"]
