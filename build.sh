#!/bin/sh

out=thomas11.github.com

rm -rf $out/*
cp -R static $out/
cp CNAME $out/

go run blog11.go
