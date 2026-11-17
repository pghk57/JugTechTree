sql_shell: 
	docker compose exec db mysql -u root -p

docker_run:
	docker compose up -d

docker_log:
	docker compose logs -f

docker_down:
	docker compose down