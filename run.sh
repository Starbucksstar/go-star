#!/bin/bash
set -e

echo "===============Wire Bean Factory============="
wire ./src/beanfactory

echo "===============Start Go-Service Web=========="
go run main.go

echo "============================================="