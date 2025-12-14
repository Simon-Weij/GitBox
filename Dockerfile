FROM node:24-slim AS frontend

WORKDIR /app

COPY ./frontend/package*.json ./
RUN npm ci

COPY ./frontend ./
RUN npm run build


FROM golang:1 AS backend

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./src ./src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./src

FROM gcr.io/distroless/base-debian13

WORKDIR /app

COPY --from=backend /app/server /app/server

COPY --from=frontend /app/build /app/dist

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
