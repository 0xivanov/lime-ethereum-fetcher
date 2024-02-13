```curl
curl -i -X POST \
  http://localhost:9090/lime/authenticate \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "alice",
    "password": "alice"
}'
```