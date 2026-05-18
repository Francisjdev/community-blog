package app

import (
	"github.com/francisjdev/community-blog/internal/config"
	"github.com/francisjdev/community-blog/internal/service"
)

type Application struct {
	Config  *config.Config
	Service *service.Services
}
