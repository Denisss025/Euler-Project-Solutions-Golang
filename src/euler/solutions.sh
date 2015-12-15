#!/bin/bash

mkdir -p "./reg-solutions/"

pkg="$1"
out="./reg-solutions/$2"
exec 3>"$out"

echo -e "package $pkg\n\nimport (" >&3

find ../ -maxdepth 1 -type d -regex '.*euler[0-9]+'  \
	| sort | sed 's/^.*euler\(.*\)$/\t_ "euler\1"/' >&3

echo ")" >&3
