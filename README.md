# go-beer-docs

### Run In Docker

```bash
# Run this and wait for everything to start up
docker-compose up --build
```

In another terminal seed the database with some beers with curl or whatever you'd like

```bash
curl -s --location --request POST 'http://127.0.0.1:8080/api/v1/beers/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "beername": "Crank Yanker IPA",
    "brewery": {
        "name": "Eddyline Brewery"
    }
}' | jq
curl -s --location --request POST 'http://127.0.0.1:8080/api/v1/beers/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "beername": "Yall Need Science",
    "brewery": {
        "name": "Cerebral Brewing"
    }
}' | jq
curl -s --location --request POST 'http://127.0.0.1:8080/api/v1/beers/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "beername": "Neon Nail",
    "brewery": {
        "name": "Our Mutual Friend"
    }
}' | jq
```