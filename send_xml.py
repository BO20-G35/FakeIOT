import requests

url = "http://127.0.0.1:8080/config"
xml_data = """
<settings>
<address>0.0.0.0</address>
<port>8081</port>
</settings>"""

r = requests.post(url,data=xml_data)
