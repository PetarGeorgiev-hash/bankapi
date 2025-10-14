#!/bin/sh

set -e

echo "run db migration"
set -a
. /app/.env
set +a
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start app"
exec "$@"