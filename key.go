package redigo_pack

import "github.com/gomodule/redigo/redis"

type KeyRds struct {
	conn redis.Conn
}

// 返回值类型
func (k *KeyRds) Type(key string) *Reply {
	return NewReply(k.conn.Do("type", key))
}

// 查找键 [*模糊查找]
func (k *KeyRds) Keys(pattern string) *Reply {
	return NewReply(k.conn.Do("keys", pattern))
}

// 随机返回一个key
func (k *KeyRds) RandomKey() *Reply {
	return NewReply(k.conn.Do("randomkey"))
}

// 判断key是否存在
func (k *KeyRds) Exists(key string) *Reply {
	return NewReply(k.conn.Do("exists", key))
}

// @param: keys — 列表
//    使用案例: Del( []string{key1, key2, xxx} )
func (k *KeyRds) Del(keys []string) *Reply {
	return NewReply(k.conn.Do("Del", redis.Args{}.AddFlat(keys)...))
}

// 重命名
func (k *KeyRds) Rename(key, newKey string) *Reply {
	return NewReply(k.conn.Do("Rename", key, newKey))
}

// 仅当newkey不存在时重命名
func (k *KeyRds) RenameNx(key, newKey string) *Reply {
	return NewReply(k.conn.Do("RenameNx", key, newKey))
}

// 序列化key
func (k *KeyRds) Dump(key string) *Reply {
	return NewReply(k.conn.Do("Dump", key))
}

// 反序列化
func (k *KeyRds) Restore(key string, ttl, serializedValue interface{}) *Reply {
	return NewReply(k.conn.Do("Restore", key, ttl, serializedValue))
}

// 同实例不同库间的键移动
func (k *KeyRds) Move(key string, db int64) *Reply {
	return NewReply(k.conn.Do("move", key, db))
}

/* --- 过期时间: expire/expireAt --- */

// 秒
func (k *KeyRds) Expire(key string, seconds int64) *Reply {
	return NewReply(k.conn.Do("expire", key, seconds))
}

// 秒
func (k *KeyRds) ExpireAt(key string, timestamp int64) *Reply {
	return NewReply(k.conn.Do("expireat", key, timestamp))
}

// 毫秒
func (k *KeyRds) PExpire(key string, milliseconds int64) *Reply {
	return NewReply(k.conn.Do("PExpire", key, milliseconds))
}

// 毫秒
func (k *KeyRds) PExpireAt(key string, millisecondsTimestamp int64) *Reply {
	return NewReply(k.conn.Do("PExpireAt", key, millisecondsTimestamp))
}

/* --- Persist | TTL  --- */

// Persist 移除过期时间
func (k *KeyRds) Persist(key string) *Reply {
	return NewReply(k.conn.Do("Persist", key))
}

// 秒
func (k *KeyRds) TTL(key string) *Reply {
	return NewReply(k.conn.Do("TTL", key))
}

// 毫秒
func (k *KeyRds) PTTL(key string) *Reply {
	return NewReply(k.conn.Do("PTTL", key))
}

/* --- 待实现 --- */
// MIGRATE \  OBJECT \  SORT \  SCAN \
