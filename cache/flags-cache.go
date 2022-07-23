package cache

import (
	"fmt"
	"manager/configs"
	"strconv"
	"time"
)
// will we need a different data type for sdk keys ? currently no model

// FlagCache is implemented by redisCache struct
type FlagCache interface {
	Set(key string, value interface{}) // set an array of
	// Get(key string) interface{}        // not needed? mgr never reads from cache
	FlushAllAsync()
}

func InitFlagCache() FlagCache {
	address := fmt.Sprintf("%s:%s", configs.GetEnvVar("REDIS_HOST"), configs.GetEnvVar("REDIS_PORT"))
	db, err := strconv.Atoi(configs.GetEnvVar("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	expires, err := time.ParseDuration(configs.GetEnvVar("SECS_TO_EXPIRE"))

	if err != nil {
		panic(err)
	}
	return NewRedisCache(address, db, expires)
	 
}

// func RefreshCache() {
// 	// Flush cache & refresh cache

// 	flagCache := InitFlagCache()
// 	flagCache.FlushAllAsync()
// 	flagCache.Set("data", fs)

// }