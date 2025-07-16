# Build Stage
FROM golang:1.23.0-alpine3.19 AS builder

RUN apk add --no-cache git

WORKDIR /temp

ENV GOPRIVATE=github.com/*/*
ARG GPR_TOKEN
RUN git config --global url.https://$GPR_TOKEN@github.com/.insteadOf https://github.com/

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

RUN go build -o app -v ./cmd/restapi/

# Final Stage
FROM gcr.io/distroless/base-debian12
LABEL Name="Users BFF" Version="v0.0.1"

WORKDIR /app

COPY --from=builder ./temp/app .

EXPOSE 8080

ENTRYPOINT ["./app"]