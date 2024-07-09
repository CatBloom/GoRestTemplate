# GoRestTemplate

## System Architecture

* Golang
* Echo
* MySQL
* SAM 
* AWS APIGateway
* AWS Lambda

## local

### docker compose
1. `cd scripts/`
1. `docker compose up`

### unit test
1. `cd scripts/`
1. `docker compose up`
1. `docker exec -it bash`
1. `gotest -v ./...`

## deploy
1. `cd scripts/`
1. `sam build`
1. `sam deploy --config-env prod --parameter-overrrides Stage=prod`