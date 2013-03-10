#!/bin/sh

rm -rf out/*
cp -R static out/
cp CNAME out/

go run blog11.go
