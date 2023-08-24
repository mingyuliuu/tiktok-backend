MYSQL_PASSWORD = PtoYbtCt

dkcompose:
	docker compose -f ./docker-compose.yml up

dkremove:
	docker compose -f ./docker-compose.yml down

clean: 
	docker container stop mysql &
	docker container stop redis &
	docker container stop rabbitmq &
	docker container stop nginx &
	docker container stop vsftpd &
	docker image rm mysql &
	docker image rm redis &
	docker image rm rabbitmq &
	docker image rm nginx &
	docker image rm vsftpd

sqlterminal:
	docker exec -it mysql bash
	mysql -u root -p${MYSQL_PASSWORD}

gin:
	go get -u github.com/gin-gonic/gin

gorm:
	go get -u gorm.io/gorm

gen:
	go get -u gorm.io/gen

gorun:
	go run *.go

.PHONY: compose stop clean sqlterminal gin gorm gen gorun