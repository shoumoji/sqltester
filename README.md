# DB (MariaDB) テスト用ツール

## 使い方

1. ```bash
   git clone https://github.com/shoumoji/sqltester.git
   ```

1. ```bash
   cd sqltester
   ```

1. ```bash
   touch schema.sql # スキーマを書く
   ```

1. ```bash
   make run # DBコンテナに入る
   ```

1. ```bash
   make stop # 終わったらDBコンテナを消す
   ```
