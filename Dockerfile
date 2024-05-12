# Reference: https://docs.docker.com/language/golang/build-images/.
# Build the application from source
FROM golang:1.22.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /service

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /service /service

EXPOSE 8080

ENTRYPOINT ["/service"]