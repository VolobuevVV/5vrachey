#!/bin/sh
# Запускаем бэкенд в фоновом режиме
/app/main &

# Запускаем nginx на переднем плане
nginx -g 'daemon off;'