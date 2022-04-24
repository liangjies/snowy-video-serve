package model

import (
	"snowy-video-serve/global"
)

type JwtBlacklist struct {
	global.SYS_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
