package server

type router struct {
	table map[string]RequestHandler

	getRequestHandlerImpl func(string) RequestHandler
}

func (r *router) GetRequestHandler(path string) RequestHandler {
	handler := r.table[path]
	if (handler == nil) && (r.getRequestHandlerImpl != nil) {
		return r.getRequestHandlerImpl(path)
	}
	return handler
}

func createDefaultRouter() *router {
	r := &router{
		table: map[string]RequestHandler{

			"api/v0/ping":         pingHandler,        // Pingpong
			"/api/v1/score":       ScoreHandler,       // Score
			"/api/v1/leaderboard": LeaderboardHandler, //leaderboard
		},
	}
	return r
}
