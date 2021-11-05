#!/usr/bin/env bash

rm -rf ../www
mkdir ../www
hugo -d ../www -v
#mkdir ../www/images
#cp -rf ../www/about_cz/index.html ../www/about
ls ../www -l
ls ../www/about_cz
cp -rf ../logo.png ../www/images
cp -rf ../favicon.png ../www
cp -rf ../favicon.png ../www/favicon.ico
pip install lxml
mkdir ../www/about
mkdir ../www/awesome-rust
python gen.py
#rm -rf ../www/about_cz
#rm -rf ../www/awesome_rust
