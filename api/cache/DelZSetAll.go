package cache

const _DelZsetAllID = "DelZsetAll"

//DelZsetAll - 範例程式
func (db *DB) DelZsetAll(key []string, key1 string) {

	strarr := []string{key1}
	db.client.EvalSha(db.scripts[_DelZsetAllID], key, strarr)

}
