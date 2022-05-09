#!/bin/sh

docker build -t yaml2json . && docker run -p 8080:8080 yaml2json