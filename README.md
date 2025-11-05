# chirpy API 

Simple Go HTTP server exposing RESTful JSON endpoints for a Twitter‑style service.

- GET  api/healthz          – readiness probe
- GET  admin/metrics        – basic metrics
- POST admin/reset          – metrics reset endpoint
- POST api/users            – add users to the database
- POST api/chirps           – validate and add chirps ("tweets") to the database

- Serves static index.html and assets
- Uses JSON decoding/encoding 
