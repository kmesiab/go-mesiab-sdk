#!/bin/bash
# run_migration.sh

# Wait for MySQL to be ready
echo "Waiting for MySQL to start..."
sleep 10  # Adjust the sleep time if necessary

# Installing goose
echo "Running Goose migrations..."
go get -u github.com/pressly/goose/v3/cmd/goose

# Run Goose migration
echo "Running Goose migrations..."
goose -dir ../migrations mysql "$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:3306)/$MYSQL_DATABASE" up
