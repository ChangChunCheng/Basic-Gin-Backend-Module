include .env

export PROJ_PATH=basic-gin-backend-module
# export GIN_MODE=release # release, debug, test

export DATE := $(shell date +%Y.%m.%d-%H%M)
export LATEST_COMMIT := $(shell git log --pretty=format:'%h' -n 1)
export BRANCH := $(shell git branch |grep -v "no branch"| grep \*|cut -d ' ' -f2)
export BUILT_ON_IP := $(shell hostname | cut -d' ' -f1 )
export BIN_DIR=./bin
export RUNTIME_VER := $(shell go version)

export BUILT_ON_OS=$(shell uname -a)
ifeq ($(BRANCH),)
BRANCH := master
endif

export COMMIT_CNT := $(shell git rev-list HEAD | wc -l | sed 's/ //g' )
export BUILD_NUMBER := ${BRANCH}-${COMMIT_CNT}
export COMPILE_LDFLAGS='-X "main.BuildDate=${DATE}" \
                -X "main.BuiltOnIP=${BUILT_ON_IP}" \
                -X "main.BuiltOnOs=${BUILT_ON_OS}" \
				-X "main.RuntimeVer=${RUNTIME_VER}" \
                -X "main.LatestCommit=${LATEST_COMMIT}" \
                -X "main.BuildNumber=${BUILD_NUMBER}"' 

build:
	@echo "Downloading go module..."
	go mod tidy
	@echo "Building go app..."
	go build -o app -ldflags $(COMPILE_LDFLAGS)

run:
	@echo "Runing app..."
	@PORT=${APP_PORT} \
		DB_HOST=${BUILT_ON_IP} \
		DB_PORT=${POSTGRESQL_PORT} \
		DB_NAME=${POSTGRESQL_DB} \
		DB_USER=${POSTGRESQL_USER} \
		DB_PASS=${POSTGRES_PASSWORD} \
		./app

lookDB:
	@echo "Going to docker to watch DB"
	@docker exec -it user-db psql -U ${POSTGRESQL_USER} ${POSTGRESQL_DB}

buildDB:
	@echo "Building docker DB..."
	@docker-compose build user-db

runDB:
	@echo "Running docker DB..."
	@docker-compose run -d user-db

restartDB:
	@echo "Restarting docker DB..."
	@docker-compose restart user-db

buildContainer:
	@echo "Building docker DB and app..."
	@docker-compose build --parallel

runOnContainer: buildContainer
	@echo "Running docker containers..."
	@docker-compose up

stop:
	@echo "Closing docker containers..."
	@docker-compose down

restart:
	@echo "Restarting docker containers..."
	@docker-compose restart

clean:
	@echo "Cleaning docker images on local and data..."
	@docker-compose down --rmi local
	sudo rm -rf ${POSTGRESQL_DATA}