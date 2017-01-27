
install:
	curl https://glide.sh/get | sh
	go get golang.org/x/tools/cmd/goimports
	go get -u github.com/jteeuwen/go-bindata/...
	glide install
	make build-assets

fmt:
	@goimports -w ./app

build-assets:
	@cd app && go-bindata -o=cmd/server/assets.go ./assets/...

build-assets-debug:
	@cd app && go-bindata -debug -o=cmd/server/assets.go ./assets/...

build-server:
	@make build-assets
	@cd ./app/cmd/server && go build 

# make にすると本番で誤って実行したとき危険ぽいので変えた方が良さそう
migrate:
	@bin/flyway -user=root -password="" -url=jdbc:mysql://localhost/webapp -locations=filesystem:db migrate

# バッチにしないと接続先管理つらそう
model:
	@xorm reverse mysql "root:@/webapp?charset=utf8" ./vendor/github.com/go-xorm/cmd/xorm/templates/goxorm app/model
	@rm app/model/schemaVersion.go

server:
	@make build-assets-debug
	@go run ./app/cmd/server/*

