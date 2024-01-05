#!/bin/bash
set -e
echo "===============Start generate swagger============="
swag init -g ../main.go
echo "===============Successfully generate=============="