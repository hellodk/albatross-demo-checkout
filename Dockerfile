FROM docker.io/library/golang:1.27-alpine AS builder
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./main.go

FROM docker.io/library/alpine:3.21
RUN adduser -D -u 65532 -H appuser
COPY --from=builder /out/app /usr/local/bin/app
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/app"]
