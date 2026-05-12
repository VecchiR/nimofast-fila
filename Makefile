.PHONY: fullstack-up stop

fullstack-up:
	docker compose up -d
	npm run --prefix frontend frontend-up

stop:
	docker compose down