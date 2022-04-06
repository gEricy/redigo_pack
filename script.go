package redigo_pack

import (
	"github.com/gomodule/redigo/redis"
)

type ScriptRds struct {
	conn redis.Conn
}

func (s *ScriptRds) DoScript(keyCount int, script string, keysAndArgs ...interface{}) *Reply {
	return NewReply(redis.NewScript(keyCount, script).Do(s.conn, keysAndArgs...))
}
