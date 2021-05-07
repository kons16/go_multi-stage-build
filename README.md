# go_multi-stage-build
goのAPIサーバーを、Dockerのマルチステージビルドで動かすサンプル。  

### ローカルでの開発
ローカルで開発は fresh を利用したホットリロードに対応しています。
```
$ docker-compose up
 Server Start!!
$ curl http://localhost:8000/json
 {"message":"hello"}
```

### マルチステージビルド
```
$ docker build -f docker/Dockerfile -t hello_app_build .
$ docker run hello_app_build -p 8000:8000
```
