# -*- coding: UTF-8 -*-

import requests
import json


Identity={
    'school_id':'20164762',
    'name':'jieli',
    'class':'1603',
    'options':'delete',
    'parameters':''
}

pictures = {  
    'img':open("/home/jieli/Pictures/2017-11-06 10-40-36 的屏幕截图.png",'rb')
}  
  

print(Identity)
json=json.dumps(Identity)
print(json)
r = requests.post('http://123.56.223.156:80/',json)
#r = requests.post('http://123.56.223.156:8080/img',Identity,files=pictures)
print(r)

