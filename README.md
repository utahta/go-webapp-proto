# GoでWebアプリを作る調査＆プロトタイプ

## 構成

|ディレクトリ|説明|
|:---|:---|
|app|アプリのソースコード|
|bin|実行バイナリ（本体 or シンボリックリンク）|
|db|データベースマイグレーションファイル|
|vendor|サードパーティー製のライブラリ|
|config|設定ファイル|

### app/

|ディレクトリ|説明|
|:---|:---|
|cmd|実行バイナリ|
|controller|コントローラー|
|model|モデル|
|lib|よしなにロジックをまとめられれば|
|assets|bindata に固めるファイルたち|
|middleware|自前の echo ミドルウェア|

### app/cmd

|ディレクトリ|説明|
|:---|:---|
|batch|バッチ|
|server|サーバ|

## パッケージ管理

[Masterminds/glide](https://github.com/Masterminds/glide)

## 設定ファイル

[spf13/viper](https://github.com/spf13/viper)

設定ファイルは、サーバ x 環境（test, dev, prod）の分だけ存在する  
それぞれ yml ファイルを準備してあげれば良さそう  
不要なのもあるはずなので、下の表であらわす

|サーバ|test|development|production|
|:---|:---|:---|:---|
|ローカル(vagrant ...etc)|◯|◯|△(= development)|
|ステージング|☓|☓|◯|
|本番|☓|☓|◯|

## マイグレーション

[flywaydb](https://flywaydb.org/)

ある程度安定していそう

## フレームワーク

[labstack/echo](https://github.com/labstack/echo)

## ORM

[go-xorm/xorm](https://github.com/go-xorm/xorm)

## セッション

[gorilla/sessions](github.com/gorilla/sessions)  
[boj/redistore](https://github.com/boj/redistore)

Redis に格納する体  
最初にある程度整えてあげる必要性を感じた

