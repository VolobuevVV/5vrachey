FROM node:20-alpine as frontend-build
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
RUN npm install vue-router@4
COPY frontend/ .
RUN npm run build

FROM golang:1.23-alpine as backend-build
WORKDIR /app/backend
COPY backend/go.mod ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM nginx:alpine as production

COPY --from=frontend-build /app/frontend/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=backend-build /app/backend /app/backend

RUN chmod +x /app/main
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

EXPOSE 80
CMD ["/app/start.sh"]