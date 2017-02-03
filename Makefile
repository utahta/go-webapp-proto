HAVE_FLYWAY := $(shell ls bin/flyway 2> /dev/null)
HAVE_GLIDE := $(shell command -v glide 2> /dev/null)

flyway:
ifndef HAVE_FLYWAY
	curl https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/4.0.3/flyway-commandline-4.0.3.tar.gz -o ./bin/flyway-4.0.3.tar.gz
	cd bin &&\
	tar zxvf flyway-4.0.3.tar.gz &&\
	ln -s flyway-4.0.3/flyway flyway &&\
	rm flyway-4.0.3.tar.gz
endif

glide:
ifndef HAVE_GLIDE
	curl https://glide.sh/get | sh
endif

install: flyway glide
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/jteeuwen/go-bindata/...
	glide install
	make build-web-assets

fmt:
	@goimports -w ./app

build-web-assets:
	@cd app/web/assets && go-bindata -o=../cmd/assets.go ./...

build-web-assets-debug:
	@cd app/web/assets && go-bindata -debug -o=../cmd/assets.go ./...

build-web-server:
	@make build-web-assets
	@cd ./app/web/cmd && go build

# make にすると本番で誤って実行したとき危険ぽいので変えた方が良さそう
migrate:
	@bin/flyway -user=root -password="" -url=jdbc:mysql://localhost/webapp -locations=filesystem:db migrate

# バッチにしないと接続先管理つらそう
model:
	@xorm reverse mysql "root:@/webapp?charset=utf8" ./vendor/github.com/go-xorm/cmd/xorm/templates/goxorm app/model
	@rm app/model/schemaVersion.go

web:
	@make build-web-assets-debug
	@go run ./app/web/cmd/*

