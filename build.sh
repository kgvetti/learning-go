#! /bin/bash

if [ ! $1 ]; then
  echo "Usage: go name_of_file (without the .go)"
  exit
fi 

6g $1.go
6l $1.6
