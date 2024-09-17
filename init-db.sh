#!/bin/bash
set -e

# Restore the database from the backup file
pg_restore -U postgres -d todos /docker-entrypoint-initdb.d/todos-backup.tar
