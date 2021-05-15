package cache

import (
	"log"
)

const _GetZsetRangeID = "GetZsetRange"

//GetZsetRange - 範例程式
func (db *DB) GetZsetRange(key []string, key1 string) (*[]ZSetResult, error) {

	strarr := []string{key1}
	res, err := db.client.EvalSha(db.scripts[_GetZsetRangeID], key, strarr).Result()
	if err != nil {
		log.Printf("GetZsetRange err : %v\n", err)
		return nil, err
	}

	reader := NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{}))
	if count > 0 {
		count = count / 2
	}
	result := make([]ZSetResult, count)

	for i := 0; i < count; i++ {
		r := &ZSetResult{}

		r.Value2 = reader.ReadString()
		r.Value1 = reader.ReadString()
		result[i] = *r
		if err != nil {
			log.Printf("GetZsetRange err : %v\n", err)
		}
	}
	return &result, nil
}
