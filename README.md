# example-caddy
This project contains a sample Caddy server which has a variety of functions.  It is meant as a concise reference for Caddy syntax, in addition to what is available in https://caddyserver.com/docs/.  



# Running the project
Needs Docker.  
While Docker is running, you can simply:  
`docker-compose up`
... at the root of this repo.

# Behaviors
There are a couple of simple behaviors of the Caddy server that are configured. 
- It is configured to listen on 443 with the internal certificate.  It will be untrusted, but it works.  (`-k` option with curl, or click through on browser)
- `https://localhost/gin/*` will *rewrite* requests below that level to the "example-gin" service (A Gin webserver with some simple endpoints)
    - Available endpoints (relative to Gin server itself, see above):  
        - `/healthz` - health endpoint
        - `/api/v1/time` - Return current system time. 
- `https://localhost/internal` - Provide a static response *from Caddy*
- `https://localhost/healthz` - Healthcheck for the Caddy service, configured as such in the docker-compose.yaml.
- `https://localhost/static/` - a static file (index.html) served by Caddy.