.DEFAULT_GOAL = help

CONTAINER_NAME := "transfer/app:latest"
APP_PORT := "9090:9090"

## Build the container.
build: 
	docker build -t $(CONTAINER_NAME) .

## Run the contaienr.
run: 
	docker run -p ${APP_PORT} $(CONTAINER_NAME)

## Starts the container.
start: build run

#------------------------------------------------------------------------
# Document file
#------------------------------------------------------------------------

# VARIABLES
NAME = Transfer app
VERSION = 0.0.1
AUTHOR=Caio Milfont

# COLORS
GREEN := $(shell tput -Txterm setaf 2)
RESET := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

## Shows help.
help:
	@echo "--------------------------------------------------------------------------------"
	@echo "Author  : ${GREEN}$(AUTHOR)${RESET}"
	@echo "Project : ${GREEN}$(NAME)${RESET}"
	@echo "Version : ${GREEN}$(VERSION)${RESET}"
	@echo "--------------------------------------------------------------------------------"
	@echo ""
	@echo "Usage:"
	@echo "  ${GREEN}make${RESET} <target>"
	@echo "Targets:"
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${GREEN}%-$(TARGET_MAX_CHAR_NUM)s${RESET} %s\n", helpCommand, helpMessage; \
		} \
	} \
{ lastLine = $$0 }' $(MAKEFILE_LIST)