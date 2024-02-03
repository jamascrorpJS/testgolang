.DEFAULT_GOAL:= run

dump:
	docker cp backup/backup_file.dump postgres:/var/lib/postgresql/data
	docker exec postgres pg_restore -U postgres -d alifacademy /var/lib/postgresql/data/backup_file.dump

run: 
	docker compose up