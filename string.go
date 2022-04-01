package redigo_pack

import "github.com/gomodule/redigo/redis"

type StringRds struct {
	conn redis.Conn
}

/* --------- SET/GET --------- */

// SET
func (s *StringRds) Set(key string, value interface{}) *Reply {
	return NewReply(s.conn.Do("set", key, value))
}

// GET
func (s *StringRds) Get(key string) *Reply {
	return NewReply(s.conn.Do("get", key))
}

// 设置并返回旧值
func (s *StringRds) GetSet(key string, value interface{}) *Reply {
	return NewReply(s.conn.Do("getset", key, value))
}

/* --------- MSET/MGET --------- */

// MSET key value [key value …]
func (s *StringRds) Mset(kv map[string]interface{}) *Reply {
	return NewReply(s.conn.Do("mset", redis.Args{}.AddFlat(kv)...))
}

// 返回多个key的值
func (s *StringRds) Mget(keys []string) *Reply {
	return NewReply(s.conn.Do("mget", redis.Args{}.AddFlat(keys)...))
}

/* --------- SET NotExists --------- */

// key不存在是在设置值
func (s *StringRds) SetNx(key string, value interface{}) *Reply {
	return NewReply(s.conn.Do("setnx", key, value))
}

// SET <key> <val> EX <second> NX
// 		不存在，设置成功，并设置超时时间
// 返回值:
//    设置成功: reply == OK    --- 不存在，新建，设置成功
//    设置失败: reply == nil   --- 已经存在
func (s *StringRds) SetNxEx(key string, value interface{}, expire int64) *Reply {
	return NewReply(s.conn.Do("set", key, value, "EX", expire, "NX"))
}

// key不存在时设置多个值
//  返回值
//      1 - 所有key设置都生效
//      0 - 任何一个key设置未生效(该key存在)
func (s *StringRds) MsetNx(kv map[string]interface{}) *Reply {
	return NewReply(s.conn.Do("MsetNx", redis.Args{}.AddFlat(kv)...))
}

/* --------- SET Expire --------- */

// 	设置key并指定生存时间
func (s *StringRds) SetEx(key string, value interface{}, seconds int64) *Reply {
	return NewReply(s.conn.Do("setex", key, seconds, value))
}

// 	设置key值并指定生存时间(毫秒)
func (s *StringRds) PsetEx(key string, value interface{}, milliseconds int64) *Reply {
	return NewReply(s.conn.Do("psetex", key, milliseconds, value))
}

/* --------- 子字符串 --------- */

// 设置子字符串
func (s *StringRds) SetRange(key string, value interface{}, offset int64) *Reply {
	return NewReply(s.conn.Do("setrange", key, offset, value))
}

// 	获取子字符串
func (s *StringRds) GetRange(key string, start, end int64) *Reply {
	return NewReply(s.conn.Do("getrange", key, start, end))
}

// 	获取字符串长度
func (s *StringRds) Strlen(key string) *Reply {
	return NewReply(s.conn.Do("Strlen", key))
}

// 	追加
func (s *StringRds) Append(key string, appendStr string) *Reply {
	return NewReply(s.conn.Do("Append", key, appendStr))
}

/* --------- 增加/减少 --------- */

// 自增: +1
func (s *StringRds) Incr(key string) *Reply {
	return NewReply(s.conn.Do("incr", key))
}

// 增加指定值: +increment
func (s *StringRds) IncrBy(key string, increment int64) *Reply {
	return NewReply(s.conn.Do("IncrBy", key, increment))
}

// 增加一个浮点值: +increment
func (s *StringRds) IncrByFloat(key string, increment float64) *Reply {
	return NewReply(s.conn.Do("IncrByfloat", key, increment))
}

// 自减: -1
func (s *StringRds) Decr(key string) *Reply {
	return NewReply(s.conn.Do("decr", key))
}

// 自减指定值: -increment
func (s *StringRds) DecrBy(key string, increment int64) *Reply {
	return NewReply(s.conn.Do("decrby", key, increment))
}
