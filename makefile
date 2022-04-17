RUN:
	
	docker-compose up
	docker exec -it postgresqldb psql -U ui_test -c "create role ui_test with login password '66556622';”
	docker exec -it postgresqldb psql -U ui_test -c "create database ui_test owner ui_test”
	docker exec -it postgresqldb psql -U ui_test -c "\l”