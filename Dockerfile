# syntax=docker/dockerfile:1
# Build the application from source
FROM golang:1-bookworm AS build-stage

ARG PACKAGE_NAME=ur_v2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /{PACKAGE_NAME}

# # Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /{PACKAGE_NAME} /{PACKAGE_NAME}

USER nonroot:nonroot

ENTRYPOINT ["/${PACKAGE_NAME}"]