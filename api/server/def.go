package server

// TokenStruct -
type TokenStruct struct {
	VendorID  int16  `json:"vendorID"`
	UserID    int64  `json:"userID"`
	ClientIP  string `json:"clientIP"`
	Timestamp int64  `json:"timestamp"`
	GameID    int    `json:"gameID"`
}

// MessageResult -
type MessageResult struct {
	Message string `json:"message"`
}

// StandardResult -
type StandardResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	// LeaderKey -
	LeaderKey = []string{"2", "test", "leaderboard"}

	_MariadbErrorMessage = map[int]string{
		// 廠商非遊戲商身分無法修改
		-30101: "Permission Denied",
		// 遊戲名稱參數錯誤
		-30102: "invalid argument 'GameName'",
		// 上下架時間參數錯誤
		-30103: "invalid argument 'GameUptime' or 'GameDowntime'",
		// 遊戲類別參數錯誤
		-30104: "invalid argument 'GameTypeID'",
		//...
	}

	_RTPTable = map[int]int{
		1: 1,
	}
)

// MariadbErrorReference -
func MariadbErrorReference(rtnCode int) string {
	s := _MariadbErrorMessage[rtnCode]
	if s == "" {
		s = "Unknown Error."
	}
	return s
}
