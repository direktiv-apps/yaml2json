#!/bin/bash

input="$1"
if=$2
of=$3
it=$4

if [[ "$it" == "base64" ]]; then
    input=`echo $1 | base64 -d --wrap=0`
elif [[ "$it" == "file" ]]; then
    input=`cat $1`
fi

echo "$input" | yq -p $2 -o $3

# cat out.tmp

# cat out.tmp
# ewoiamVucyI6ICJnZXJrZSIKCn0=
# "data": "{\r\n\"jens\": \"gerke\"\r\n\r\n}",
#echo "converting form $if to $of"

#ho $1 $2 $3 $4p
