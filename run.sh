#!/bin/sh
dockerize --wait http://SMS_database
cd /SMS
go run main.go