from xml.dom import minidom
import os
import sys
import threading
from threading import Thread
import time

if len(sys.argv) != 2:
	print("need path to xml file")
	sys.exit(2)

def thread_parse_xml():
	mydoc = minidom.parse(sys.argv[1])
	items = mydoc.getElementsByTagName('lolz')

parse_thread = Thread(target=thread_parse_xml,daemon=True)
parse_thread.start()

time.sleep(5)

if threading.active_count() == 2:
	print("BOMB")
	sys.exit(1)

print("OK")
sys.exit(0)