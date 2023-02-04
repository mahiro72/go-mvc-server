# go-mvc-server

## Getting Started

トップディレクトリに`.env`ファイルを作成し、`.sample.env`を参考に環境変数を用意します。

上記を作成後以下のコマンドを実行し、docker環境を立ち上げます。

```
make up
```

docker環境が立ち上がったら、http://localhost:8080/health にアクセスし、サーバーが起動しているか確認してください。

okとでれば問題ないです。

またこのタイミングで`docker/db/dev/initdb.d`が実行されdummy dataがdbに入ります。

## ディレクトリ

```
|--.gitignore
|--.sample.env
|--app
|  |--.air.toml
|  |--cmd                // ここからサーバーを起動します
|  |--db
|  |  |--migrations      // migrationsファイルがまとめてあります
|  |--go.mod
|  |--go.sum
|  |--pkg
|  |  |--controller      // modelとviewを呼び出します、またcmd配下のハンドラーと対応しています
|  |  |--model           // ドメインオブジェクトの管理やpersistanceを呼び出しデータを保存します
|  |  |--persistence     // データの永続化を責務とする
|  |  |--view            // Goの構造体をJSON形式に変更する
|--docker                // dev,prodのDockerfileがまとめてあります
|--docker-compose.yml
|--Makefile              // プロジェクトに関するコマンドがまとめてあります
|--README.md
|--scripts               // Makefileから実行するスクリプトがあります
```


## Tips

### makeについて

以下のコマンドのどちらかを打つとmakeのhelpを確認できます

```
make or make help
```

## うまく起動しなくなった場合

下記のコマンドでvolumeやimageを再ビルドできます。

```
make re
```

