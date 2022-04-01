package redigo_pack

import "github.com/gomodule/redigo/redis"

type SetRds struct {
	conn redis.Conn
}

// 添加元素 SAdd key member [member …]
func (s *SetRds) SAdd(key string, fileds []interface{}) *Reply {
	return NewReply(s.conn.Do("SAdd", redis.Args{}.Add(key).AddFlat(fileds)...))
}

// 移除指定的元素 SRem key member [member …]
func (s *SetRds) SRem(key string, members []interface{}) *Reply {
	return NewReply(s.conn.Do("SRem", redis.Args{}.Add(key).AddFlat(members)...))
}

// 集合元素个数
func (s *SetRds) SCard(key string) *Reply {
	return NewReply(s.conn.Do("SCard", key))
}

// 返回集合中所有成员
func (s *SetRds) SMembers(key string) *Reply {
	return NewReply(s.conn.Do("SMembers", key))
}

// 判断元素是否是集合成员
func (s *SetRds) SIsMember(key string, member interface{}) *Reply {
	return NewReply(s.conn.Do("SIsMember", key, member))
}

// 随机返回一个或多个元素 SRANDMEMBER key [count]
func (s *SetRds) SRandMember(key string, count ...int64) *Reply {
	if len(count) > 0 {
		return NewReply(s.conn.Do("SRandMember", key, count[0]))
	}
	return NewReply(s.conn.Do("SRandMember", key))
}

// 随机返回并移除一个元素
func (s *SetRds) SPop(key string) *Reply {
	return NewReply(s.conn.Do("SPop", key))
}

/* --------- 两个\多个集合间的操作 --------- */

// 将元素member从集合srcKey移至另一个集合dstKey
func (s *SetRds) SMove(srcKey, dstKey string, member interface{}) *Reply {
	return NewReply(s.conn.Do("SMove", srcKey, dstKey, member))
}

// 返回一或多个集合的差集 SDiff key [key …]
func (s *SetRds) SDiff(keys []string) *Reply {
	return NewReply(s.conn.Do("SDiff", redis.Args{}.AddFlat(keys)...))
}

// 将一或多个集合的差集保存至另一集合(dstKey) SDIFFSTORE destination key [key …]
func (s *SetRds) SDiffStore(dstKey string, keys []string) *Reply {
	return NewReply(s.conn.Do("SDiffstore", redis.Args{}.Add(dstKey).AddFlat(keys)...))
}

// 一个或多个集合的交集 SInter key [key …]
func (s *SetRds) SInter(keys []string) *Reply {
	return NewReply(s.conn.Do("SInter", redis.Args{}.AddFlat(keys)...))
}

// 将keys的集合的并集 写入到 dstKey中   SINTERSTORE destination key [key …]
func (s *SetRds) SInterStore(dstKey string, keys []string) *Reply {
	return NewReply(s.conn.Do("SInterstore", redis.Args{}.Add(dstKey).AddFlat(keys)...))
}

// 返回集合的并集 SUnion key [key …]
func (s *SetRds) SUnion(keys []string) *Reply {
	return NewReply(s.conn.Do("SUnion", redis.Args{}.AddFlat(keys)...))
}

// 将 keys 的集合的并集 写入到 dstKey 中  SUNIONSTORE destination key [key …]
func (s *SetRds) SUnionStore(dstKey string, keys []string) *Reply {
	return NewReply(s.conn.Do("SUnionstore", redis.Args{}.Add(dstKey).AddFlat(keys)...))
}

/* --------- SScan --------- */

// SScan key cursor [MATCH pattern] [COUNT count]
func (s *SetRds) SScan(key string, cursor int64, pattern string, count int64) *Reply {

	var args []interface{}

	args = append(args, key)
	args = append(args, cursor)

	if pattern != "" {
		args = append(args, "MATCH")
		args = append(args, pattern)
	}

	if count > 0 {
		args = append(args, "COUNT")
		args = append(args, count)
	}

	return NewReply(s.conn.Do("SScan", args...))
}
