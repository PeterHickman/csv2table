#!/bin/sh

BINARY='/usr/local/bin'
APP=csv2table

echo "Building $APP"
go build -ldflags="-s -w" $APP.go

echo "Installing $APP to $BINARY"
install $APP $BINARY

echo "Removing the build"
rm $APP
