.PHONY: help

PATH_TO_ENV := .env
ENV_EXISTS := $(or $(and $(wildcard $(PATH_TO_ENV)),1),0)

help: ## Show command list
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

chkenv: ## Check if env file exists
ifneq (${ENV_EXISTS}, 1)
	@ echo "Error: No ".env" file" && exit 1
endif

vendor: ## Generate vendor folder
	@ go mod vendor

build: chkenv vendor ## Build the API
	@ docker-compose -f devops/docker-compose.yml --project-name cartesian build --no-cache

run: chkenv vendor ## Run the API
	@ docker-compose -f devops/docker-compose.yml --project-name cartesian up --build --remove-orphans --detach

logs: ## Watch API logs
	@ docker-compose -f devops/docker-compose.yml --project-name cartesian logs --follow api

mysql_logs: ## Watch MySQL logs
	@ docker-compose -f devops/docker-compose.yml --project-name cartesian logs --follow mysql

redis_logs: ## Watch Redis logs
	@ docker-compose -f devops/docker-compose.yml --project-name cartesian logs --follow redis

redis: ## Attach to redis-cli
	@ docker-compose -f devops/docker-compose.yml --project-name cartesian exec redis redis-cli

test: vendor ## Test the API  
	@ docker-compose -f devops/docker-compose-test.yml --project-name cartesian_test --env-file .env-example \
		up --build --remove-orphans --detach
	@ docker-compose -f devops/docker-compose-test.yml --project-name cartesian_test \
		logs --follow api-test