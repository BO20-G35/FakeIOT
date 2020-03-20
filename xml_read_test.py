from xml.dom import minidom
import os
import sys
import threading
from threading import Thread
import time


def thread_parse_xml():
	mydoc = minidom.parse('/home/tobias/go/src/FakeIOT/bomb.xml')
	items = mydoc.getElementsByTagName('lolz')
	


parse_thread = Thread(target=thread_parse_xml,daemon=True)
parse_thread.start()

time.sleep(5)

if threading.active_count() == 2:
	print("BOMB")
	sys.exit(1)

print("OK")
sys.exit(0)