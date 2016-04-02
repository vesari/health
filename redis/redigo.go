package redis

import (
	"errors"

	redigo "github.com/garyburd/redigo/redis"
)

type connectionProvider interface {
	getConnection(network, addr string) (redigo.Conn, error)
}

type defaultConnectionProvider struct{}

func (c defaultConnectionProvider) getConnection(network, addr string) (redigo.Conn, error) {
	return redigo.Dial(network, addr)
}

// RedigoRedis implements the interface Redis with redigo package
type RedigoRedis struct {
	network string
	addr    string

	provider connectionProvider
}

// NewRedigo create a new redigoRedis with specified network and addr
func NewRedigo(network, addr string) RedigoRedis {
	return RedigoRedis{
		network:  network,
		addr:     addr,
		provider: defaultConnectionProvider{},
	}
}

// GetVersion return the redis version using the reigo package
func (r RedigoRedis) GetVersion() (string, error) {

	conn, err := r.provider.getConnection(r.network, r.addr)

	if err != nil {
		return "", err
	}

	defer conn.Close()

	data, err := redigo.String(conn.Do("INFO"))

	if err != nil {
		return "", err
	}

	info := parseInfo(data)

	version, ok := info["redis_version"]

	if !ok {
		return "", errors.New("redis_version is not present in INFO response")
	}

	return version, nil
}
