
#!/bin/bash
echo -e "Testing..."
curl -s http://127.0.0.1:8080/api/v1/beers/ | jq
echo
curl -s http://127.0.0.1:8080/api/v1/beers/2 | jq
echo
curl -s http://127.0.0.1:8080/api/v1/breweries/ | jq
echo
curl -s http://127.0.0.1:8080/api/v1/breweries/1 | jq
echo