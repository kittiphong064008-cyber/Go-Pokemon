# ---------- Stage 1: build ----------
ARG GO_VERSION=1.25
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /src

# copy go.mod/go.sum ก่อน เพื่อให้ layer ของ dependency ถูก cache
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# lib/pq เป็น pure Go จึงปิด CGO ได้ -> binary static รันบน image เล็กๆ ได้
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath -ldflags="-s -w" \
    -o /out/app ./cmd/api

# ---------- Stage 2: run ----------
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /out/app /app

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app"]