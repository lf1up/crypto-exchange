project_name = crypto-api
image_name = app
db_image_name = db
db_volume_name = db-volume

help: ## This help dialog.
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run app.go

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

up: ## Run the project in a local container
	make up-silent
	make shell

build: ## Generate docker image
	docker build -f Dockerfile.app -t ${image_name}:$(project_name)-v1 .
	docker build -f Dockerfile.db -t ${db_image_name}:$(project_name)-v1 .

build-no-cache: ## Generate docker image with no cache
	docker build -f Dockerfile.app --no-cache -t ${image_name}:$(project_name)-v1
	docker build -f Dockerfile.db --no-cache -t ${db_image_name}:$(project_name)-v1

up-silent: ## Run local container in background
	make delete-container-if-exist
	docker run -d -p 5432:5432 --name $(project_name)-${db_image_name} -v $(project_name)-${db_volume_name}:/var/lib/postgresql/data ${db_image_name}:$(project_name)-v1
	docker create -p 3000:3000 --name $(project_name)-${image_name} ${image_name}:$(project_name)-v1 ./app
	make create-network
	docker start $(project_name)-${image_name}

up-silent-prefork: ## Run local container in background with prefork
	make delete-container-if-exist
	docker run -d -p 5432:5432 --name $(project_name)-${db_image_name} ${db_image_name}:$(project_name)-v1
	docker create -p 3000:3000 --name $(project_name)-${image_name} ${image_name}:$(project_name)-v1 ./app -prod
	make create-network
	docker start $(project_name)-${image_name}

delete-container-if-exist: ## Delete container if it exists
	docker network rm $(project_name)-network || true
	docker stop $(project_name)-${image_name} || true && docker rm $(project_name)-${image_name} || true
	docker stop $(project_name)-${db_image_name} || true && docker rm $(project_name)-${db_image_name} || true

create-network: ## Create a network
	docker network create $(project_name)-network
	docker network connect $(project_name)-network $(project_name)-${db_image_name}
	docker network connect $(project_name)-network $(project_name)-${image_name}

shell: ## Run interactive shell in the container
	docker exec -it $(project_name)-${image_name} /bin/sh

stop: ## Stop the container
	docker stop $(project_name)-${image_name}
	docker stop $(project_name)-${db_image_name}

start: ## Start the container
	docker start $(project_name)-${image_name}
	docker start $(project_name)-${db_image_name}

purge-db-volume: ## Purge the database volume
	docker stop $(project_name)-${db_image_name}
	docker rm $(project_name)-${db_image_name}
	docker volume rm $(project_name)-${db_volume_name}
