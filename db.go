package redigo_pack

import "github.com/gomodule/redigo/redis"

type DbRds struct {
	conn redis.Conn
}

// switch db
func (d *DbRds) Select(db int64) *Reply {
	return NewReply(d.conn.Do("select", db))
}
