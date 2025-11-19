# chirpy 

Simple experimental Go HTTP server exposing RESTful JSON endpoints for a Twitter‑style service.

## Features

- Create/Delete short text posts ("chirps")
- View recent chirps
- Register and login users 
- Password hashing with Argon2id
- Authentication and authorization with JWT tokens with 
- Refresh or revoke tokens 
- Basic metrics and readiness 
- Webhook for a fictional external third-party payment service  

- Serves static index.html and assets
- Uses JSON decoding/encoding 
- PostgreSQL database integration 

## Avaliable endpoints

- GET  api/healthz          – readiness probe
- GET  admin/metrics        – basic metrics
- POST admin/reset          – Database and metrics reset endpoint

- POST api/users            – add users to the database
- POST api/login            - authenticate and login user 
- PUT api/users             – user update 

- POST api/chirps           – validate and add chirps ("tweets") to the database
- GET api/chirps            – get all chirps from the database with query parameters for Author ID and ASC/DESC order  
- GET api/chirps/{chirpID}  - get chirp by ID 

- POST api/refresh          - refresh JWT token 
- POST api/revoke           - revoke JWT token 

- POST api/polka/webhooks   - webhook endpoint for third-party payment service to confirm subscription activation

## Dependencies 

- [github.com/alexedwards/argon2id](https://github.com/alexedwards/argon2id)
Used for secure password hashing with the Argon2id key-derivation function.

- [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt/v5)
Used to generate and validate JSON Web Tokens (JWT) for authentication.

- [github.com/google/uuid](https://github.com/google/uuid)
Used to generate RFC 4122 compliant UUIDs for identifiers.

- [github.com/joho/godotenv](https://github.com/joho/godotenv)
Used to load configuration from a .env file into environment variables.

- [github.com/lib/pq](https://github.com/lib/pq)
PostgreSQL driver for Go’s database/sql package.

- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)
Additional cryptographic primitives used alongside argon2id and JWT.

- [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys)
Low-level system call helpers used by other dependencies.
