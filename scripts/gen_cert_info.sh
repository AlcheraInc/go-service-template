#!/bin/bash
#
#   Generate 'formed' input for openssl certificate generation ...
#

echo "KR"
echo "Gyeonggi"
echo "Seongnam"
echo "personal"
echo "developer"
echo ${1:-"$(hostname)"}
echo ${2:-"unknown@mail.com"}
