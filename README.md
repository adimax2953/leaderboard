# leaderboard

題目

Leader BoardDesign and implement (with unit tests) leader board server using Go programming language.
Criteria:
Theleader board server has 2 APIs, please follow API example to implement:
A RESTful API to receive gaming score from a client (with its client ID).
The leader board service should reset leader board every 10 minutes.
This assignment is multi-server design.
Please feel free to use any external libs if needed.It is also free to use following external storage including:
Relational database (MySQL, PostgreSQL, SQLite)Cache storage (Redis, Memcached)You do not need to consider auth.
Please implement reasonable constrains and error handling of these 2 APIs

如發現Redis腳本遺失，請通知我，謝謝。


測試步驟

1.首先開個終端機 go run app.go

2.看到 apitest listening on address :9090 log 的時候到score_test.go內

3.第24行的上方有個run test 給他點下去，他會提交分數到 redis內

4.第161行的上方有個run test給他點下去，可以獲取排行榜前十名的資訊(理論上來說應會把這個功能做在另個Worker上)

5.server.go的69行有一個抓每小時內的每十分鐘的排程去做排行榜重置

當然作法很多種，有興趣我們可以一起討論。
