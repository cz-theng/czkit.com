#!/usr/bin/env bash

rm -rf ../www/*
hugo -d ../www -v
#mkdir ../www/images
cp -rf ../www/about_cz/index.html ../www/about
cp -rf ../logo.png ../www/images
cp -rf ../favicon.png ../www
cp -rf ../favicon.png ../www/favicon.ico
pip install lxml
python gen.py
rm -rf ../abount_cz
rm -rf ../awesome_rust
