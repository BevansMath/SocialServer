#!/bin/sh

echo "Enter your message"
git add .
read MESSAGE
git commit -m "${MESSAGE}"
git push  origin master

