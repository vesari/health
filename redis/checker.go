package redis

import "github.com/dimiro1/health"

// Redis is a interface used to abstract the access of the Version string
type Redis interface {
	GetVersion() (string, error)
}

// Checker is a checker that check a given redis
type Checker struct {
	Redis Redis
}

// NewChecker returns a new redis.Checker
func NewChecker(redis Redis) Checker {
	return Checker{Redis: redis}
}

// Check obtain the version string from redis info command
func (c Checker) Check() health.Health {
	health := health.NewHealth()
	health.Unknown()

	version, err := c.Redis.GetVersion()

	if err != nil {
		health.Down().AddInfo("error", err.Error())
		return health
	}

	health.Up().AddInfo("version", version)

	return health
}
