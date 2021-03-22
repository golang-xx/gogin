#!/bin/bash

git add .
git commit -m 'first time'

echo 'Push to old-origin main'
git push old-origin main
echo 'Push to origin main'
git push origin main

