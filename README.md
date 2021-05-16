# leaderboard

如發現Redis腳本遺失，請通知我，謝謝。


測試步驟

1.首先開個終端機 go run app.go

2.看到 apitest listening on address :9090 log 的時候到score_test.go內

3.第24行的上方有個run test 給他點下去，他會提交分數到 redis內

4.第161行的上方有個run test給他點下去，可以獲取排行榜前十名的資訊

5.server.go的69行有一個抓每小時內的每十分鐘的排程去做排行榜重置

當然作法很多種，有興趣我們可以一起討論。
