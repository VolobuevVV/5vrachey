#!/bin/sh

cd /app
./backend/main &
sleep 3
nginx -g 'daemon off;'