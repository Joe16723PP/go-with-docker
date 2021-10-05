#!/usr/bin/env bash
echo 'Running migrations...'
/go-with-docker/migrate up > /dev/null 2>&1 &

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Start application...'
/go-with-docker/app