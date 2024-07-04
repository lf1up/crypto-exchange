# Crypto Exchange API Example

## IDE Development

### Visual Studio Code

Use the following plugins, in this boilerplate project:
- Name: Go
  - ID: golang.go
  - Description: Rich Go language support for Visual Studio Code
  - Version: 0.29.0
  - Editor: Go Team at Google
  - Link to Marketplace to VS: https://marketplace.visualstudio.com/items?itemName=golang.Go

## Development

### Fill the *.env* file

```
FASTFOREX_API_KEY=<PUT_YOUR_KEY_HERE>
POSTGRES_USER=<PUT_YOUR_PG_USERNAME_HERE>
POSTGRES_PASSWORD=<PUT_YOUR_PG_PASSWORD_HERE>
POSTGRES_DB=<PUT_YOUR_PG_DB_NAME_HERE>
POSTGRES_HOST=127.0.0.1
```

### Start the application in dev mode (hot code reloading)


```bash
make build
make up-dev
```

### Use local containers

```
# Shows all commands
make help

# Clean packages
make clean-packages

# Generate go.mod & go.sum files
make requirements

# Generate docker images
make build

# Generate docker images with no cache
make build-no-cache

# Run the project and SSH into the main container
make up

# Run local containers in background
make up-silent

# Run local containers in background with prefork
make up-silent-prefork

# Stop containers
make stop

# Start containers
make start

# Purge the database volume
make purge-db-volume
```

## Production

```bash
# See the Makefile script section for build info
make up-silent-prefork
```

## A quick intorduction
Access the API at http://localhost:3000/api/v1/

* http://localhost:3000/api/v1/pairs `[GET]` Fetch list of all avaliabie currency pairs.
* http://localhost:3000/api/v1/pairs/*<PAIR_NAME>* `[GET]` Fetch details about currency *<PAIR_NAME>*.
* http://localhost:3000/api/v1/convert `[POST]` Do the value amount conversion between two currency pairs.

See the `app.go` for more info.