#!/bin/bash
set -e

echo "===============Wire Bean Factory============="
wire ../src/**

echo "===============Start Go-Service Web=========="
go run ../main.go

echo "============================================="
