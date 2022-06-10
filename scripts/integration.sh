#!/bin/bash

mkdir -p artifacts
set -e

# Get the container ip for the service
ip=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' users-service)
echo "container ip: $ip"

echo "creating user"
# Create a user
./bin/userctl create -a $ip --first-name=testfn --last-name=testln --nickname=testnn \
--password=testpw --email=test@email.com --country=BU > artifacts/create.json

id=$(cat artifacts/create.json | jq '.id')
echo "storing id: $id"

# Check JSON response
cat artifacts/create.json | jq -e 'has("id")'
cat artifacts/create.json | jq -e 'has("created_at")'
cat artifacts/create.json | jq -e 'has("updated_at")'
cat artifacts/create.json | jq -e '.first_name | contains("testfn")'
cat artifacts/create.json | jq -e '.last_name | contains("testln")'
cat artifacts/create.json | jq -e '.nickname | contains("testnn")'
cat artifacts/create.json | jq -e '.password | contains("testpw")'
cat artifacts/create.json | jq -e '.email | contains("test@email.com")'
cat artifacts/create.json | jq -e '.country | contains("BU")'

echo "updating user"
# Update a user
./bin/userctl update -a $ip --id=$(eval echo $id) --first-name=testfn2 \
--last-name=testln2 --nickname=testnn2 --password=testpw2 \
--email=test@email.com2 --country=AU > artifacts/update.json

# Check JSON response
cat artifacts/update.json | jq -e 'has("id")'
cat artifacts/update.json | jq -e 'has("created_at")'
cat artifacts/update.json | jq -e 'has("updated_at")'
cat artifacts/update.json | jq -e '.first_name | contains("testfn2")'
cat artifacts/update.json | jq -e '.last_name | contains("testln2")'
cat artifacts/update.json | jq -e '.nickname | contains("testnn2")'
cat artifacts/update.json | jq -e '.password | contains("testpw2")'
cat artifacts/update.json | jq -e '.email | contains("test@email.com2")'
cat artifacts/update.json | jq -e '.country | contains("AU")'

echo "listing users"
# List Users
./bin/userctl list -a $ip --country=UA > artifacts/list.json

# Check JSON response
cat artifacts/list.json | jq -e 'has("id")'
cat artifacts/list.json | jq -e 'has("created_at")'
cat artifacts/list.json | jq -e 'has("updated_at")'
cat artifacts/list.json | jq -e '.first_name | contains("john")'
cat artifacts/list.json | jq -e '.last_name | contains("smith")'
cat artifacts/list.json | jq -e '.nickname | contains("john-smith")'
cat artifacts/list.json | jq -e '.password | contains("john-test-pw")'
cat artifacts/list.json | jq -e '.email | contains("john@test.com")'
cat artifacts/list.json | jq -e '.country | contains("UA")'

echo "deleting user"
# Delete Users
./bin/userctl delete -a $ip --id=$(eval echo $id)
