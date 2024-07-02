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

### Start the application 


```bash
go run app.go
```

### Use local container

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

Use API at http://localhost:3000/api/v1/
