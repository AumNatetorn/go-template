curl -k -X POST http://localhost:8080/template \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'

curl http://localhost:8080/health