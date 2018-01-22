import httplib
 
httpClient = None
 
try:
    httpClient = httplib.HTTPConnection('127.0.0.1', 8080, timeout=30)
    httpClient.request('GET', '/')
 
    response = httpClient.getresponse()
    print response.status
    print response.reason
    print response.read()
except Exception, e:
    print e
finally:
    if httpClient:
        httpClient.close()