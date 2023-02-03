# go-mvc-server

## Getting Started

トップディレクトリに`.env`ファイルを作成し、`.sample.env`を参考に環境変数を用意します。

上記を作成後以下のコマンドを実行し、docker環境を立ち上げます。

```
make up
```

docker環境が立ち上がったら、http://localhost:8080/health にアクセスし、サーバーが起動しているか確認してください。

okとでれば問題ないです。

## Check

動作確認は以下を参考にしてください。



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

