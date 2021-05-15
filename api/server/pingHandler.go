package server

import (
	"log"

	"github.com/valyala/fasthttp"
)

func pingHandler(ctx *fasthttp.RequestCtx) {
	method := ctx.Method()

	switch string(method) {
	case "GET":
		log.Println(ctx.Request.URI().String())
		{
			ctx.Success("charset=utf8", []byte("pong"))
			return
		}

	case "POST":
		log.Println(ctx.Request.URI().String())
		{
			ctx.Success("charset=utf8", []byte("pong"))
			return
		}

	case "PUT":
		log.Println(ctx.Request.URI().String())
		{
			ctx.Success("charset=utf8", []byte("pong"))
			return
		}

	case "PATCH":
		log.Println(ctx.Request.URI().String())
		{
			ctx.Success("charset=utf8", []byte("pong"))
			return
		}

	case "DELETE":
		log.Println(ctx.Request.URI().String())
		{
			ctx.Success("charset=utf8", []byte("pong"))
			return
		}

	default:
		ctx.Error("Method Not Allowed", fasthttp.StatusMethodNotAllowed)
		return
	}
}
