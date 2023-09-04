deploy:
	docker-compose build
	docker-compose up

verify_db:
	echo "docker exec -it backend-assignments_postgres_1 bash"
	echo 'psql "postgresql://postgres:changeme@postgres:5432/postgres"'
	echo "select * from log_entry;"