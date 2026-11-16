sql_shell: 
	docker compose exec db mysql -u root -p

sql_run:
	docker compose up -d