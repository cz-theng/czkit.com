#!/usr/bin/python2.7 
from lxml.etree import HTML
from lxml import etree, html
from lxml.html.clean import Cleaner
import os

def parse_blog(src,dst):
	src_path = os.path.abspath("../www/"+src+"/index.html")
	print("src:",src_path)
	in_file = open(src_path,"r")
	html_text = in_file.read()
	dom= html.fromstring(html_text)
	pc_divs=dom.xpath('.//div[@class="post-copyright"]')
	pm_divs=dom.xpath('.//div[@class="post-meta"]')
	pc_divs[0].getparent().remove(pc_divs[0])
	pm_divs[0].getparent().remove(pm_divs[0])
	dst_path = os.path.abspath("../www/"+dst+"/index.html")
	print("dst:",dst_path)
	out_file = open(dst_path,"w")
	out_file.write(html.tostring(dom))

def main():
	parse_blog("about_cz", "about")
	parse_blog("awesome_rust","awesome-rust")

main()
