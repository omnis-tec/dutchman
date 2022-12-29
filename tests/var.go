package tests

import (
	"github.com/rendau/dop/adapters/cache/mem"
	dopDbPg "github.com/rendau/dop/adapters/db/pg"
	dopLoggerZap "github.com/rendau/dop/adapters/logger/zap"
	"github.com/rendau/dutchman/internal/adapters/repo/pg"
	"github.com/rendau/dutchman/internal/domain/core"
	"github.com/rendau/dutchman/internal/domain/usecases"
)

var (
	app = struct {
		lg    *dopLoggerZap.St
		cache *mem.St
		db    *dopDbPg.St
		repo  *pg.St
		core  *core.St
		ucs   *usecases.St
	}{}
)
