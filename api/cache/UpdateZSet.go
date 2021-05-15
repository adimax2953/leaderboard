package cache

import (
	"log"
)

// ZSetResult -
type ZSetResult struct {
	Value1 string
	Value2 string
}

const _UpdateZsetID = "UpdateZset"

//UpdateZset - 範例程式
func (db *DB) UpdateZset(key []string, key1, value1, value2 string) (*ZSetResult, error) {

	strarr := []string{key1, value1, value2}
	res, err := db.client.EvalSha(db.scripts[_UpdateZsetID], key, strarr).Result()
	if err != nil {
		log.Printf("UpdateZset err : %v\n", err)
		return nil, err
	}

	result := &ZSetResult{}
	reader := NewRedisArrayReplyReader(res.([]interface{}))
	result.Value1 = reader.ReadString()
	if result.Value1 == "-1" {
		log.Printf("UpdateZset Value1 err : %v\n", err)
		return nil, err
	}
	result.Value2 = reader.ReadString()
	if result.Value2 == "-1" {
		log.Printf("UpdateZset Value2 err : %v\n", err)
		return nil, err
	}
	return result, nil
}
