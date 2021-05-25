## Dapper Labs - Auth Service

*Contact: Email nick@dhupia.org for any questions. Please note that none of the **credentials** in the git repo are for production use and are provided for convenience only. You must replace these **credentials** with your own **unless** using for demo/eval purposes only.*

**Following files contain security credentials**
   * `server.key and server.crt are used for https - ssl/tls.`
   * `The .env file contains postgres db credentials and jwt secret key.`


**Simply run:**
- `git clone git@github.com:axiomzen/cc_NickDhupia_BackendAPI.git` clone the repo
- `docker-compose up --build --remove-orphans` to bring up the app cleanly
- `docker-compose down` to shutdown and cleanup the app

**Sample Tests**
1. `curl -v -k -d '{"email":"1@db.com", "password":"password","firstName":"f1", "lastName":"l1"}' -H 'Content-Type: application/json' -X POST https://localhost:8081/signup`
2. `curl -v -k -d '{"email":"2@db.com", "password":"password","firstName":"f2", "lastName":"l2"}' -H 'Content-Type: application/json' -X POST https://localhost:8081/signup`
3. `curl -v -k -d '{"email":"3@db.com", "password":"password","firstName":"f3", "lastName":"l3"}' -H 'Content-Type: application/json' -X POST https://localhost:8081/signup`
4. `curl -v -k -d '{"email":"1@db.com", "password":"password"}' -X POST https://localhost:8081/login`
5. `curl -v -k -d '{"email":"2@db.com", "password":"password"}' -X POST https://localhost:8081/login`
6. `curl -v -k -H'x-authentication-token:#TOKEN' -X GET https://localhost:8081/users`
7. `curl -v -k -d '{"firstName":"TomFirstName", "lastName":"JerryLastName"}' -H'x-authentication-token:#TOKEN' -X PUT https://localhost:8081/users`