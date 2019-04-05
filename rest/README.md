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

- [ ] ユーザーごとに「取引」（金額と商品説明からなる情報）を登録することができるサービスです。
      現状は特になんの制限もなく、何件でも取引登録可能です。

      これに対し、デイリーで取引登録可能な上限金額を設けたいという要件が出てきたという状況です。


- [ ] total > limitになったら Post: /transactonsはstatus StatusPaymentRequired(402)を返却