# go-beer-docs

## Run In Docker

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
curl -s --location --request POST 'http://127.0.0.1:8080/api/v1/beers/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "beername": "Epic Day Double IPA",
    "brewery": {
        "name": "Eddyline Brewery"
    }
}' | jq
```

## Adding New Routes

When new routes are added in Go, the React client lib needs to be generated again.

In the project root run `swag init`.  Then generate the lib in the client directory.

```bash
cd client/
bash gen_client_fetch_lib.sh
```

TODO — automate this.
