# cc-server
chat-connectのサーバー。

## URL
- APIサーバー：[http://localhost:8001]()
- mockサーバー：[http://localhost:8002]()

## 環境構築
1.コンテナを起動
- APP_ENV=development：air
- APP_ENV=production：go build
```
docker compose up -d --build
```
## API
1.サーバーを起動
```
docker container exec -it cc-server-api-1 go run api/main.go
```

## Batch
2.バッチを実行
```
docker container exec -it cc-server-batch-1 go run batch/main.go --command=example
```

## Swagger
1.Swaggerのビルド
```
docker container exec -it cc-server-api-1 swag init --dir=api --output=swagger
```
2.Swaggerのmackサーバーを起動
```
docker container exec -it cc-server-swagger-1 prism mock ./swagger/swagger.yaml --port=8002 --host=0.0.0.0
```

## DI
1.API
```
docker container exec -it cc-server-api-1 wire api/di/wire.go
```

2.Batch
```
docker container exec -it cc-server-api-1 wire batch/di/wire.go
``
