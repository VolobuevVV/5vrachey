# Frontend build stage - используем Node.js 20
FROM node:20-alpine as frontend-build
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ .
RUN npm run build

# Backend build stage - используем Go 1.23
FROM golang:1.23-alpine as backend-build
WORKDIR /app/backend
COPY backend/go.mod ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Production stage - используем nginx для фронтенда и запускаем бэкенд
FROM nginx:alpine as production

# Копируем собранный фронтенд в nginx
COPY --from=frontend-build /app/frontend/dist /usr/share/nginx/html

# Копируем nginx конфиг
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Копируем бэкенд бинарник
COPY --from=backend-build /app/backend/main /app/main

# Делаем бэкенд исполняемым
RUN chmod +x /app/main

# Запускаем и бэкенд и nginx через скрипт
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

EXPOSE 80
CMD ["/app/start.sh"]