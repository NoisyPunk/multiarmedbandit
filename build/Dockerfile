FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

# Dependencies
COPY ["go.mod", "go.sum", "./"]
COPY . .
RUN go mod download

# Build
COPY . ./
RUN go build -o ./bin/rotator cmd/main.go

#
FROM alpine:3.9 AS runner
#
COPY --from=builder /usr/local/src/bin /

COPY configs/* /rotator_config.yaml
COPY migrations/* /

CMD ["/rotator"]

