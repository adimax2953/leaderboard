package server

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func testService(t *testing.T, service *Service) *Service {
	c, _ := NewClient(&Option{
		URL:                       "http://localhost:9090",
		Name:                      "testAPI",
		MaxConnsPerHost:           512,
		MaxIdemponentCallAttempts: 5,

		PostScoreURI:      "/api/v1/score",
		GetLeaderBoardURI: "/api/v1/leaderboard",
	})
	return c
}

func TestPostScore(t *testing.T) {
	svc := testService(t, nil)

	tests := []struct {
		name        string
		clientid    string
		c           *Service
		wantDataOut *ScoreResult
		wantErr     bool
		testFixture ScoreData
	}{
		{
			name:        "提交分數",
			clientid:    "test1",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 5000.1},
		},
		{
			name:        "提交分數",
			clientid:    "test2",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 99.1},
		},
		{
			name:        "提交分數",
			clientid:    "test3",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 33.1},
		},
		{
			name:        "提交分數",
			clientid:    "test4",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 107.1},
		},
		{
			name:        "提交分數",
			clientid:    "test5",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 55.1},
		},
		{
			name:        "提交分數",
			clientid:    "test6",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 688.1},
		},
		{
			name:        "提交分數",
			clientid:    "test7",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 990.1},
		},
		{
			name:        "提交分數",
			clientid:    "test8",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 77.1},
		},
		{
			name:        "提交分數",
			clientid:    "test9",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 87.1},
		},
		{
			name:        "提交分數",
			clientid:    "test10",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 1000.1},
		},
		{
			name:        "提交分數",
			clientid:    "test11",
			c:           svc,
			wantErr:     false,
			testFixture: ScoreData{Score: 0.1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataOut, err := tt.c.PostScore(tt.clientid, tt.testFixture)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.PostScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDataOut, tt.wantDataOut) {
				t.Errorf("Service.PostScore() = %+v, want %+v", gotDataOut, tt.wantDataOut)
			}
		})
	}
}

// PostScore  提交分數
func (c *Service) PostScore(clientid string, dataIn ScoreData) (dataOut *ScoreResult, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.URL + c.PostScoreURI)
	req.Header.Set("clientId", clientid)

	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod("POST")
	body, err := json.Marshal(dataIn)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	if err := c.Client.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()
	if resp.StatusCode() != 200 {
		return nil, err
	}

	dataOut = &ScoreResult{}
	if err := json.Unmarshal(bodyBytes, dataOut); err != nil {
		return nil, err
	}
	return dataOut, nil
}

func TestGetLeaderBoard(t *testing.T) {
	svc := testService(t, nil)

	tests := []struct {
		name        string
		c           *Service
		wantDataOut *ScoreResult
		wantErr     bool
	}{
		{
			name:    "取得排行榜",
			c:       svc,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataOut, err := tt.c.GetLeaderboard()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetLeaderboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDataOut, tt.wantDataOut) {
				t.Errorf("Service.GetLeaderboard() = %+v, want %+v", gotDataOut, tt.wantDataOut)
			}
		})
	}
}

// GetLeaderboard 取得排行榜
func (c *Service) GetLeaderboard() (dataOut *LeaderboardResult, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.URL + c.GetLeaderBoardURI)
	req.Header.Add("X-ServerID", c.Name)
	req.Header.SetContentType("application/json; charset=utf-8")
	req.Header.SetMethod("GET")
	if err := c.Client.Do(req, resp); err != nil {
		return nil, err
	}

	bodyBytes := resp.Body()
	if resp.StatusCode() != 200 {
		return nil, err
	}

	dataOut = &LeaderboardResult{}
	if err := json.Unmarshal(bodyBytes, dataOut); err != nil {
		return nil, err
	}
	return dataOut, nil
}
