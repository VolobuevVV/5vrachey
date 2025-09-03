# Frontend build stage - используем Node.js 20
FROM node:20-alpine as frontend-build
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

# Backend build stage - ОБНОВЛЕНО до Go 1.23
FROM golang:1.23-alpine as backend-build
WORKDIR /app/backend
COPY backend/go.mod ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Production stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=frontend-build /app/frontend/dist ./frontend/dist
COPY --from=backend-build /app/backend/main .
EXPOSE 8080
CMD ["./main"]