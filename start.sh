#!/bin/sh

cd /app
./main &
sleep 3
nginx -g 'daemon off;'