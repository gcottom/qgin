package qgin

import (
	"context"

	"github.com/gcottom/qgin/middleware"
	"github.com/gin-gonic/gin"
)

type Config struct {
	UseContextMW       bool
	UseLoggingMW       bool
	UseRequestIDMW     bool
	InjectRequestIDCTX bool
	LogRequestID       bool
	ProdMode           bool
}

var activeConfig Config

func NewGinEngine(ctx *context.Context, cfg *Config) *gin.Engine {
	if cfg != nil && cfg.ProdMode {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()

	if cfg != nil {
		setActiveConfig(cfg)
		if cfg.UseRequestIDMW {
			engine.Use(middleware.RequestIDMiddleware(middleware.UUIDGenerator))
		}
		if cfg.UseContextMW {
			engine.Use(middleware.ContextMiddleware(*ctx))
		}
		if cfg.UseLoggingMW {
			engine.Use(middleware.LoggingMiddleware())
		}

	}
	engine.Use(gin.Recovery())
	return engine
}

func setActiveConfig(cfg *Config) {
	activeConfig = *cfg
	middleware.InjectRequestIDCTX = cfg.InjectRequestIDCTX
}

func GetActiveConfig() Config {
	return activeConfig
}
