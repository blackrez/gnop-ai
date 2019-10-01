#/usr/bin/env sh
curl -vvvv -X POST http://localhost:8080/upload \
  -F "file=@./dog.jpg" \
  -H "Content-Type: multipart/form-data"