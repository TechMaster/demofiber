package cookie_session

import (
	"demofiber/model"
	"encoding/gob"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

var Sess *session.Store

func InitSession() *redis.Storage {
	redisDB := redis.New(redis.Config{
		Host:     "localhost",
		Port:     6379,
		Username: "",
		Password: "123",
		Database: 0,
		Reset:    false,
	})

	Sess = session.New(session.Config{
		//KeyLookup: "cookie:TechMaster",

		CookieName: "TechMaster",
		Expiration: 24 * time.Hour,
		Storage:    redisDB,
	})
	gob.Register(model.Authenticate{})
	return redisDB
}
