# Gunakan image golang sebagai base untuk membuild
FROM golang:1.23 as builder

# Setel working directory dalam container
WORKDIR /app

# Salin semua file dari direktori lokal ke dalam direktori kerja container
COPY . .

# Build aplikasi Go
# RUN go mod init example.com/app || true  # Init module jika belum ada go.mod
# RUN go mod tidy                          # Download dependencies jika ada
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go             # Build file main.go

# Gunakan image yang lebih ringan untuk menjalankan aplikasi
FROM alpine:latest

# Setel working directory untuk runtime
WORKDIR /app

# Salin executable dari tahap builder ke tahap runtime
COPY --from=builder /app/main /app/
COPY config.yaml /app/config.yaml

# Jalankan executable
CMD ["/app/main"]
