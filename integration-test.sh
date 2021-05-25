curl -v -k -d '{"email":"1a@db.com", "password":"password","firstName":"f1", "lastName":"l1"}' -H 'Content-Type: application/json' -X POST https://localhost:8081/signup
curl -v -k -d '{"email":"2a@db.com", "password":"password","firstName":"f2", "lastName":"l2"}' -H 'Content-Type: application/json' -X POST https://localhost:8081/signup
curl -v -k -d '{"email":"3a@db.com", "password":"password","firstName":"f3", "lastName":"l3"}' -H 'Content-Type: application/json' -X POST https://localhost:8081/signup
curl -k -d '{"email":"1a@db.com", "password":"password"}' -X POST https://localhost:8081/login
TOKEN='add-token-here'
curl -k -H'x-authentication-token:${TOKEN}' -X GET https://localhost:8081/users
curl -k -d '{"firstName":"TomFirstName", "lastName":"JerryLastName"}' -H'x-authentication-token:${TOKEN}' -X PUT https://0.0.0.0:8081/users