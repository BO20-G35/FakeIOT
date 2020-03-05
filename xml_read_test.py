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


#tiden det tar å lese inn xml filen linje 9/10
#vil ta lang tid hvis de har lagd en xml bombe
#vi kan ta tiden før den bli lest inn så etter 
#hvis tiden er si 10 sec så har de klart det
#bakdel er at de må ikke lage en for stor bombe 