package global

import (
	"gorm.io/gorm"
	"star/src/config"
	"time"
)

var (
	GlobalMysqlClient *gorm.DB
	MysqlConfig       *config.MysqlConfiguration
)

const TokenExpireDuration = time.Hour * 24

var CustomSecret = []byte("To be a good cow horse")
