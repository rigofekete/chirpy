# chirpy API 

## Small Go HTTP server for simple endpoints and JSON handling.

- GET  api/healthz          – readiness probe
- GET  admin/metrics        – basic metrics
- POST admin/reset          – metrics reset endpoint
- POST api/validate_chirp   – validates short text messages (<= 140 chars) and replaces profane words

- Serves static index.html and assets
- Uses JSON decoding/encoding 
