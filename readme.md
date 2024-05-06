## 概要

UR賃貸住宅から部屋のデータを収集して、DBに保存するアプリ (Golang rewrite version)

https://www.ur-net.go.jp/chintai/

## 動かし方

### 動作環境

Docker が動く環境

### データ収取を実行

```
git clone https://github.com/shibu1x/ur_v2.git
cd ur_v2
docker compose run --rm app
```

### データを確認

データベース接続クライアント アプリを使用し下記に接続して確認する
```
host: 127.0.0.1
user: dev
pass: dev
db:   ur_v2
net:  tcp
```

### 終了

```
docker compose down
```

## その他

### なんでこんなことしてるの？

- 人気高いので、空きが出てもすぐに埋まってしまう
- 埋まってしまうと部屋の情報が見れなくなってしまい、どのような条件だったのかも分からない
- そのため埋まる前に情報を集めたかった

### goroutines

使いたかったが、短時間に多くのリクエストを送るとブロックされるので使えなかった
