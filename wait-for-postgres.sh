#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

# Ensure DB_PASSWORD is set
if [ -z "$DB_PASSWORD" ]; then
  echo "DB_PASSWORD is not set. Exiting."
  exit 1
fi

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd
