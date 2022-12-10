OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

.PHONY: start
start:
	@echo -e "$(OK_COLOR)==> Starting application...$(NO_COLOR)"
	@docker network create lambda-local || exit 0
	@docker-compose up -d
	@cd src && make build
	@echo -e "$(OK_COLOR)==> Run dynamodb-admin $(NO_COLOR)"
	@cd src && sam local start-api --docker-network lambda-local

.PHONY: stop
stop:
	@echo -e "$(WARN_COLOR)==> Stopping application...$(NO_COLOR)"
	@docker-compose down
	@docker network rm lambda-local
	@echo -e "$(OK_COLOR)==> Stopped application...$(NO_COLOR)"

.PHONY: run
run:
	@cd src && reflex -c reflex.conf
