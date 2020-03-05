from xml.dom import minidom
import os
import psutil


process = psutil.Process(os.getpid())
print(process.memory_info().rss) #11MB

mydoc = minidom.parse('bomb.xml')
items = mydoc.getElementsByTagName('lolz')


print(len(items))
print(items)

process = psutil.Process(os.getpid())
print(process.memory_info().rss)  # 42MB