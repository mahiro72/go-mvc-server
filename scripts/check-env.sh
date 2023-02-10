#!/bin/bash

ENV_FILE=".env"

if [ ! -e $ENV_FILE ]; then
  echo "error: .env file not found"
  exit 1
fi

echo ".env file found"
exit 0
