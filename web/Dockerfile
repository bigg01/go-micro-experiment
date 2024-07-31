FROM docker.io/golang:1.22-alpine as build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 go build -o server .

FROM gcr.io/distroless/static-debian12

COPY --from=build /app/server  /
CMD ["./server"]