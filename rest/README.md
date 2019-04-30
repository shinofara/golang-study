# rest

## 初期準備

docker for macのインストール

## 初期構築

```
$ docker-compose up -d
```

ログの確認

```
$ docker-compose logs -f app
```

## Goの開発準備

`export GO111MODULE=on`
を実行

今回はgo modを使います。

## 動作確認

```
$ curl -H "apikey: secure-api-key-1" http://127.0.0.1:8888/
Hello World
```

## テスト

```
$ go test ./...
```

## やる事

取引登録できるAPI

- [ ] total > limitになったら Post: /transactonsはstatus StatusPaymentRequired(402)を返却
- [ ]