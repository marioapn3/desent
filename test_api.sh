#!/bin/bash

BASE="http://localhost:1234"

echo "========================================="
echo " LEVEL 1: Ping"
echo "========================================="
echo "GET /ping"
curl -s "$BASE/ping" | jq .
echo ""

echo "========================================="
echo " LEVEL 2: Echo"
echo "========================================="
echo "POST /echo"
curl -s -X POST "$BASE/echo" \
  -H "Content-Type: application/json" \
  -d '{"message":"hello","from":"tester"}' | jq .
echo ""

echo "========================================="
echo " LEVEL 3: CRUD - Create & Read"
echo "========================================="
echo "POST /books (Book 1)"
curl -s -X POST "$BASE/books" \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Go Programming","author":"John"}' | jq .
echo ""

echo "POST /books (Book 2)"
curl -s -X POST "$BASE/books" \
  -H "Content-Type: application/json" \
  -d '{"id":"2","title":"Clean Architecture","author":"Uncle Bob"}' | jq .
echo ""

echo "POST /books (Book 3)"
curl -s -X POST "$BASE/books" \
  -H "Content-Type: application/json" \
  -d '{"id":"3","title":"Refactoring","author":"John"}' | jq .
echo ""

echo "--- Get Auth Token for protected routes ---"
TOKEN=$(curl -s -X POST "$BASE/auth/token" | jq -r '.token')
echo "Token: ${TOKEN:0:30}..."
echo ""

echo "GET /books"
curl -s "$BASE/books" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "GET /books/1"
curl -s "$BASE/books/1" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "========================================="
echo " LEVEL 4: CRUD - Update & Delete"
echo "========================================="
echo "PUT /books/1"
curl -s -X PUT "$BASE/books/1" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"id":"1","title":"Go Programming Updated","author":"John Doe"}' | jq .
echo ""

echo "GET /books/1 (verify update)"
curl -s "$BASE/books/1" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "DELETE /books/3"
curl -s -X DELETE "$BASE/books/3" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "GET /books (verify delete)"
curl -s "$BASE/books" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "========================================="
echo " LEVEL 5: Auth Guard"
echo "========================================="
echo "POST /auth/token"
curl -s -X POST "$BASE/auth/token" | jq .
echo ""

echo "GET /books (tanpa token - harus 401)"
curl -s "$BASE/books" | jq .
echo ""

echo "GET /books (dengan token invalid - harus 401)"
curl -s "$BASE/books" -H "Authorization: Bearer invalidtoken123" | jq .
echo ""

echo "GET /books (dengan token valid - harus 200)"
curl -s "$BASE/books" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "========================================="
echo " LEVEL 6: Search & Paginate"
echo "========================================="
echo "GET /books?author=John Doe"
curl -s "$BASE/books?author=John%20Doe" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "GET /books?author=Uncle Bob"
curl -s "$BASE/books?author=Uncle%20Bob" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "GET /books?page=1&limit=1"
curl -s "$BASE/books?page=1&limit=1" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "GET /books?page=2&limit=1"
curl -s "$BASE/books?page=2&limit=1" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "========================================="
echo " LEVEL 7: Error Handling"
echo "========================================="
echo "POST /books (invalid - tanpa title)"
curl -s -X POST "$BASE/books" \
  -H "Content-Type: application/json" \
  -d '{"id":"99","title":"","author":"Someone"}' | jq .
echo ""

echo "POST /books (invalid - tanpa author)"
curl -s -X POST "$BASE/books" \
  -H "Content-Type: application/json" \
  -d '{"id":"99","title":"Some Book","author":""}' | jq .
echo ""

echo "GET /books/999 (not found)"
curl -s "$BASE/books/999" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "PUT /books/999 (not found)"
curl -s -X PUT "$BASE/books/999" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"id":"999","title":"Ghost","author":"Nobody"}' | jq .
echo ""

echo "DELETE /books/999 (not found)"
curl -s -X DELETE "$BASE/books/999" -H "Authorization: Bearer $TOKEN" | jq .
echo ""

echo "========================================="
echo " LEVEL 8: Boss - Speed Run (All Endpoints)"
echo "========================================="
echo "Semua level 1-7 sudah dijalankan di atas."
echo "========================================="
echo " DONE"
echo "========================================="
