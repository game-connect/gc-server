# cc-server
chat-connectのサーバー。

## URL
- APIサーバー：[http://localhost:8001]()
- mockサーバー：[http://localhost:9000]()
## 環境構築
1.コンテナを起動
- APP_ENV=development：air
- APP_ENV=production：go build
```
docker compose up -d --build
```
2.Swaggerのビルド
```
docker container exec -it cc-server-api-1 swag init --output=swagger
```
3.Swaggerのmackサーバーを起動
```
docker container exec -it cc-server-swagger-1 prism mock ./swagger/swagger.yaml --port=9000 --host=0.0.0.0
```

## テスト
1.model
```
docker container exec -it cc-server-api-1 go test -v ./domain/model
```
