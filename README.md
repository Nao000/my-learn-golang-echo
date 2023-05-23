# my-learn-golang-echo

Go言語とEchoとNext.jsの勉強用リポジトリです

## Go

https://go.dev/doc/

特に Accessing databases を見る

## Go Echo

echo guide
https://echo.labstack.com/guide/

## Next.js

https://nextjs.org/docs

Go言語 Echo で Template を表示するのが難しそうだったので、フロントエンドの何かを探しました。

Next.js が SSG, SSR, CSR をページごとに選択できるらしく、途中で選択ができるっぽくて良さそうなので選びました。

なので Next.js も学び中です。

Next.js プロジェクト例

https://github.com/vercel/next.js/tree/canary/examples

## DB にテストデータを流し込むとき

1. `docker exec -it db bash`
2. `cd /home/test_db/`
3. `mysql < employees.sql -h localhost -u root -ppassword`
   1. 50秒ぐらいかかる
4. おわり