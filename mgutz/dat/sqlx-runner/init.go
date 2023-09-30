package runner

import (
	"time"

	"github.com/unix-world/go-modules/mgutz/dat"
	"github.com/unix-world/go-modules/mgutz/dat/kvs"
	"github.com/unix-world/go-modules/mgutz/dat/postgres"
	"github.com/mgutz/logxi/v1"
)

var logger log.Logger

// LogQueriesThreshold is the threshold for logging "slow" queries
var LogQueriesThreshold time.Duration

func init() {
	dat.Dialect = postgres.New()
	logger = log.New("dat:sqlx")
}

// Cache caches query results.
var Cache kvs.KeyValueStore

// SetCache sets this runner's cache. The default cache is in-memory
// based. See cache.MemoryKeyValueStore.
func SetCache(store kvs.KeyValueStore) {
	Cache = store
}
