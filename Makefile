include .env

export PROJ_PATH=basic-gin-backend-module
# export GIN_MODE=release # release, debug, test

export DATE := $(shell date +%Y.%m.%d-%H%M)
export LATEST_COMMIT := $(shell git log --pretty=format:'%h' -n 1)
export BRANCH := $(shell git branch |grep -v "no branch"| grep \*|cut -d ' ' -f2)
export BUILT_ON_IP := $(shell hostname -I | cut -d' ' -f1 )
export BIN_DIR=./bin
export RUNTIME_VER := $(shell go version)

export BUILT_ON_OS=$(shell uname -a)
ifeq ($(BRANCH),)
BRANCH := master
endif

export COMMIT_CNT := $(shell git rev-list HEAD | wc -l | sed 's/ //g' )
export BUILD_NUMBER := ${BRANCH}-${COMMIT_CNT}
export COMPILE_LDFLAGS="main.BuildDate=${DATE}" \
                "main.BuiltOnIP=${BUILT_ON_IP}" \
                "main.BuiltOnOs=${BUILT_ON_OS}" \
				"main.RuntimeVer=${RUNTIME_VER}" \
                "main.LatestCommit=${LATEST_COMMIT}" \
                "main.BuildNumber=${BUILD_NUMBER}"  

run:
	@PORT=${APP_PORT} \
		DB_HOST=${POSTGRESQL_HOST} \
		DB_PORT=${POSTGRESQL_PORT} \
		DB_NAME=${POSTGRESQL_DB} \
		DB_USER=${POSTGRESQL_USER} \
		DB_PASS=${POSTGRES_PASSWORD} \
		go run main.go $(COMPILE_LDFLAGS)

lookDB:
	@docker exec -it user-db psql -U ${POSTGRESQL_USER} ${POSTGRESQL_DB}

build:
	@docker-compose build --parallel

runDB:
	@docker-compose up -d

stopDB:
	@docker-compose down

restartDB:
	docker-compose restart

cleanDB:
	@docker-compose down --rmi local
	sudo rm -rf ${POSTGRESQL_DATA}

removeContainer:
	@docker-compose down --rmi local