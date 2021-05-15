package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/textproto"
	"strconv"

	//	"log"

	"github.com/valyala/fasthttp"
)

// ScoreData -
type ScoreData struct {
	Score float64 `json:"score"`
}

// ScoreResult -
type ScoreResult struct {
	Status string `json:"status"`
}

// ScoreHandler -
func ScoreHandler(ctx *fasthttp.RequestCtx) {
	method := ctx.Method()
	switch string(method) {
	case "POST":
		log.Println(string(ctx.PostBody()))
		{

			args := &ScoreData{}
			result := &ScoreResult{}
			if err := json.Unmarshal(ctx.PostBody(), &args); err != nil {
				result.Status = "fail"
			} else {
				score := fmt.Sprintf("%f", args.Score)
				clientid := string(ctx.Request.Header.PeekBytes([]byte(textproto.CanonicalMIMEHeaderKey("ClientId"))))
				_, err := redis.UpdateZset(LeaderKey, "score", score, clientid)
				if err == nil {
					result.Status = "ok"
				}
			}

			body, err := json.Marshal(result)
			if err == nil {
				ctx.Success("application/json; charset=utf8", body)
			}
			return
		}
	default:
		ctx.Error("Method Not Allowed", fasthttp.StatusMethodNotAllowed)
		return
	}
}

// ScoreData -
type TopPlayer struct {
	ClientID string  `json:"clientId"`
	Score    float64 `json:"score"`
}

// LeaderboardResult -
type LeaderboardResult struct {
	TopPlayers []TopPlayer `json:"topPlayers"`
}

// LeaderboardHandler -
func LeaderboardHandler(ctx *fasthttp.RequestCtx) {
	method := ctx.Method()
	switch string(method) {
	case "GET":
		log.Println(string(ctx.PostBody()))
		{
			result := &LeaderboardResult{}

			playersarr, err := redis.GetZsetRange(LeaderKey, "score")
			if err != nil {
				ctx.Error("application/json; charset=utf8", fasthttp.StatusNoContent)
				return
			}
			arr := *playersarr
			count := len(arr)
			if count >= 10 {
				count = 10
			}
			for i := 0; i < count; i++ {

				score, _ := strconv.ParseFloat(arr[i].Value1, 64)
				pack := TopPlayer{
					ClientID: arr[i].Value2,
					Score:    score,
				}
				result.TopPlayers = append(result.TopPlayers, pack)
			}
			log.Printf("%v", result)
			body, err := json.Marshal(result)
			if err == nil {
				ctx.Success("application/json; charset=utf8", body)
			}
			return
		}

	default:
		ctx.Error("Method Not Allowed", fasthttp.StatusMethodNotAllowed)
		return
	}
}
