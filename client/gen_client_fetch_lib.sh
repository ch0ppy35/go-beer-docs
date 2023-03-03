#!/bin/bash

echo "Generating the client lib..."
npx openapi --input ../docs/swagger.json --output ./src/generated
echo "Adding in api base, this is hacky"
sed -i.bak "s/BASE: '',/BASE: 'http:\/\/127.0.0.1:8080\/api\/v1',/" src/generated/core/OpenAPI.ts
