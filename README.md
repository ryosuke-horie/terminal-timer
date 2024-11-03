# terminal-timer(golang習作用)

## 概要

ターミナル上で動作するタイマーです。
分と秒を指定して、指定した時間が経過すると終わったことがわかるシンプルなタイマー。

## 使い方

```bash
# 5秒のタイマー
go run main.go 0 5

# 1分30秒のタイマー
go run main.go 1 30
```

## 自分のWSL環境のコマンドとして利用する

```bash
# ビルド >> timerという名前で実行ファイルが生成される
go build -o timer main.go
# パスを通す
sudo mv timer /usr/local/bin
# 使い方: timer [分] [秒]
timer 0 5
```
