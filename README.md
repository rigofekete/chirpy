# Chirpy API – A small Go HTTP server for simple endpoints and JSON handling.

- GET  api/healthz – readiness probe
- GET  admin/metrics – basic metrics
- POST admin/reset – development reset endpoint
- POST api/validate_chirp – validates short text messages (<= 140 chars)
- Serves static index.html and assets
- Uses JSON decoding/encoding 
