CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
MOCKGEN=${BINDIR}/mockgen_${GOVER}
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.49.0
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=ozonProjectmodule/cmd/bot
REPORT_PACKAGE=gitlab.ozon.dev/bl4ckv0id/project-base/cmd/reportService

dev:
	go run ${PACKAGE} -devel

report:
	go run ${REPORT_PACKAGE} -devel

prod:
	mkdir -p logs/data
	go run ${PACKAGE} 2>&1 | tee logs/data/log.txt

.PHONY: kafka
kafka:
	cd kafka && docker compose up

.PHONY: tracing
tracing:
	cd tracing && docker compose up

.PHONY: tracing
tracing-sudo:
	cd tracing && sudo docker compose up

.PHONY: metrics
metrics:
	mkdir -p metrics/data
	chmod -R 777 metrics/data
	cd metrics && docker compose up

.PHONY: metrics
metrics-sudo:
	mkdir -p metrics/data
	sudo chmod -R 777 metrics/data
	cd metrics && sudo docker compose up

pull:
	sudo docker pull prom/prometheus
	sudo docker pull grafana/grafana-oss
	sudo docker pull ozonru/file.d:latest-linux-amd64
	sudo docker pull elasticsearch:7.17.6
	sudo docker pull graylog/graylog:4.3
	sudo docker pull jaegertracing/all-in-one:1.18

db:
	docker compose -f docker-compose-db.yaml up

run-dev:
	docker compose up -f docker-compose.yml -f docker-compose.dev.yml --build

all: format build test lint

build: bindir
	echo ${BINDIR}/bot ${PACKAGE}
	go build -o ${BINDIR}/bot ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

generate: install-mockgen
	${MOCKGEN} \
		-source=internal/model/messages/incoming_msg.go \
		-destination=internal/mocks/messages/messages_mocks.go

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-mockgen: bindir
	test -f ${MOCKGEN} || \
		(GOBIN=${BINDIR} go install github.com/golang/mock/mockgen@v1.6.0 && \
		mv ${BINDIR}/mockgen ${MOCKGEN})

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

docker-run:
	docker compose up
