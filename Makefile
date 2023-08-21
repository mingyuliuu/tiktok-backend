CONTAINER_NAME = gosql
MYSQL_NAME = tiktok_backend
MYSQL_PORT = 3306
MYSQL_USERNAME = root
MYSQL_PASSWORD = PtoYbtCt

# Download Image If Not Exists Then Build Mysql Container (database_name = ${MYSQL_NAME}, default_user_name = root, default_user_password = ${MYSQL_PASSWORD})
composedb:
ifeq ($(shell docker images -q mysql 2> /dev/null),)  
	docker pull mysql
endif
	docker run --name ${CONTAINER_NAME} -d -e MYSQL_DATABASE=${MYSQL_NAME} -e MYSQL_ROOT_PASSWORD=${MYSQL_PASSWORD} -p ${MYSQL_PORT}:${MYSQL_PORT} mysql 

# Stop Then Delete Mysql Container
dropdb:
	docker stop ${CONTAINER_NAME}
	docker rm ${CONTAINER_NAME}

# Open sql terminal (required to manual inpit the second line)
sqlterminal:
	docker exec -it ${CONTAINER_NAME} bash
	mysql -u root -p${MYSQL_PASSWORD}

# Clean Environment
clean: dropdb 
	docker rmi mysql

# Download GIN
gin:
	go get -u github.com/gin-gonic/gin

# Download GORM
gorm:
	go get -u gorm.io/gorm

# Download Gen
gen:
	go get -u gorm.io/gen

# Run Project
run:
	go run *.go

.PHONY: composedb dropdb sqlterminal clean gin gorm run