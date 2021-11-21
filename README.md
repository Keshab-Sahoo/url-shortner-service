# url-shortner-service

Endpoints added

1.Health endpoint - to check if server is running or not
http://localhost:8080/health
HTTP method: GET

2.URL shortner endpoint - returns a short URL
http://localhost:8080/short-url?longURL=http://google.com/1346461234567890123456789/get/who_is_who
HTTP Method: GET
Sample JSON Response:
---------------------
{
    "originalURL": "http://google.com/1346461234567890123456789/get/who_is_who",
    "shortURL": "http://localhost:8080/2TrAnj1lRpP^"
}

3. Redirect endpoint - It will redirect the request to the original url
http://localhost:8080/2TrAnj1lRpP^ will redirect to http://google.com/1346461234567890123456789/get/who_is_who


Steps to execute application locally:
-------------------------------------
Note: Go should be installed local machine

1.To run test cases > go test
2.To run url-shortner-service > go run server.go or >./url-shortner-service.exe
3.Then we can test the above end endpoints

I faced some issue to run docker comand due to some network policy restriction.
So I have built an executable, which can be used to start the application anywhere.
