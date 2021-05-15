package cache

import (
	"errors"
	"sync"

	"github.com/go-redis/redis"
)

// DBOption -
type DBOption struct {
	Addr                  string
	Password              string
	DB                    int
	PoolSize              int
	RedisScriptDefinition string
	RedisScriptDB         int
}

// DB -
type DB struct {
	client               *redis.Client
	sRedisClientSyncOnce sync.Once
	scripts              map[string]string
}

// NewDB -
func NewDB(option *DBOption) (*DB, error) {
	if option == nil {
		return nil, errors.New("'option' cannot be nil")
	}
	db := &DB{}
	var err error

	// 建立連線
	db.sRedisClientSyncOnce.Do(func() {
		db.client = redis.NewClient(
			&redis.Options{
				Addr:     option.Addr,
				Password: option.Password,
				DB:       option.DB,
				PoolSize: option.PoolSize,
			})
	})
	_, err = db.client.Ping().Result()
	if err != nil {
		panic(err)
	}

	// get lua script sha
	if option.RedisScriptDefinition != "" {
		cacheScriptSet := NewDBScriptSet()
		err = cacheScriptSet.Load(db.client,
			option.RedisScriptDefinition,
			option.RedisScriptDB,
		)
		if err != nil {
			panic(err)
		}
		db.scripts = cacheScriptSet.container
	} else {
		db.scripts = map[string]string{}
	}

	return db, nil
}

//
// DBScriptSet - Load the lua script sha
//
var (
	loadSctiptTemplate = `redis.pcall('SELECT', ARGV[1])
return redis.call('HGETALL', KEYS[1])`
)

// DBScriptSet -
type DBScriptSet struct {
	container map[string]string
}

// NewDBScriptSet -
func NewDBScriptSet() *DBScriptSet {
	return &DBScriptSet{
		container: make(map[string]string),
	}
}

// SetScriptID -
func (s *DBScriptSet) SetScriptID(name string, signature string) {
	s.container[name] = signature
}

// Load -
func (s *DBScriptSet) Load(client *redis.Client, key string, db int) error {
	if client == nil {
		return errors.New("client is invalid")
	}
	res, err := client.Eval(loadSctiptTemplate, []string{key}, db).Result()
	if err != nil {
		return err
	}
	if v, ok := res.([]interface{}); ok {
		count := len(v)
		for i := 0; i < count; i = i + 2 {
			key, value := v[i], v[i+1]
			s.container[key.(string)] = value.(string)
		}
	}

	return nil
}

// Range -
func (s *DBScriptSet) Range(fn func(name string, signature string)) {
	for i, v := range s.container {
		fn(i, v)
	}
}

// Stop - shutdown
func (db *DB) Stop() {
	db.client.Close()
}
