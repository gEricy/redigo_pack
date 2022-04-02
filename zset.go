package redigo_pack

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type ZsetRds struct {
	conn redis.Conn
}

/* --------- 增 --------- */

// ZAdd key score member [[score member] [score member] …]

// map[member]score
func (z *ZsetRds) ZAddFromMap(key string, member2scoreMap map[interface{}]interface{}) *Reply {

	// score,member至少出现一对
	if len(member2scoreMap) == 0 {
		return NewReply(nil, fmt.Errorf("invalid param, len(member2scoreMap)=%d", len(member2scoreMap)))
	}

	var args []interface{}

	args = append(args, key)

	for member, score := range member2scoreMap {
		args = append(args, score)
		args = append(args, member)
	}

	return NewReply(z.conn.Do("ZAdd", args...))
}

/* --------- 删 --------- */

// @param:   members []string — 列表
// @example: ZRem(key, []string{"member1", "member2", "member3"})
func (z *ZsetRds) ZRem(key string, members []string) *Reply {
	return NewReply(z.conn.Do("ZRem", redis.Args{}.Add(key).AddFlat(members)...))
}

// ZRemRangeByRank key start stop
//     移除有序集 key 中，指定排名(Rank)区间内的所有成员
func (z *ZsetRds) ZRemRangeByRank(key string, strat, stop interface{}) *Reply {

	var args []interface{}

	args = append(args, key)
	args = append(args, strat)
	args = append(args, stop)

	return NewReply(z.conn.Do("ZRemRangeByRank", args...))
}

// ZRemRangeByScore key min max
//     移除有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员
func (z *ZsetRds) ZRemRangeByScore(key string, min, max interface{}) *Reply {

	var args []interface{}

	args = append(args, key)
	args = append(args, min)
	args = append(args, max)

	return NewReply(z.conn.Do("ZRemRangeByScore", args...))
}

/* --------- 改 --------- */

// 为成员 member 的 score 值加上增量 increment
func (z *ZsetRds) ZIncrBy(key string, increment, member interface{}) *Reply {
	return NewReply(z.conn.Do("ZIncrBy", key, increment, member))
}

/* --------- 查 --------- */

// @return 成员 member 的总个数
func (z *ZsetRds) ZCard(key string) *Reply {
	return NewReply(z.conn.Do("ZCard", key))
}

// ------ 查: 成员 member 的score

// @return 成员 member 的 score 值
func (z *ZsetRds) ZScore(key string, member interface{}) *Reply {
	return NewReply(z.conn.Do("ZScore", key, member))
}

// ------ 查: 成员 member 的排名

// @return 成员 member 的排名
// 说明: 其中有序集成员按 score 值递增(从小到大)顺序排列，排名以 0 为底，也就是说， score 值最小的成员排名为 0
func (z *ZsetRds) ZRank(key string, member interface{}) *Reply {
	return NewReply(z.conn.Do("ZRank", key, member))
}

func (z *ZsetRds) ZRevRank(key string, member interface{}) *Reply {
	return NewReply(z.conn.Do("ZRevRank", key, member))
}

// ------ 查: 按照 score

// ZCount key min max
// @return  score 值在 [min, max] 之间的成员的数量
func (z *ZsetRds) ZCount(key string, min, max interface{}) *Reply {
	return NewReply(z.conn.Do("ZCount", key, min, max))
}

// @return 返回有序集 key 中，所有 score 处于 {min, max}之间 的成员
//     ZRangeByScore key min max [WITHSCORES] [LIMIT offset count]
func (z *ZsetRds) ZRangeByScore(key string, min, max interface{}, withscore bool, limit ...interface{}) *Reply {

	var args []interface{}

	args = append(args, key)
	args = append(args, min)
	args = append(args, max)

	if withscore {
		args = append(args, "WITHSCORES")
	}

	if len(limit) == 2 {
		args = append(args, "LIMIT")
		args = append(args, limit[0]) // offset
		args = append(args, limit[1]) // count
	}

	return NewReply(z.conn.Do("ZRangeByScore", args...))
}

func (z *ZsetRds) ZRevRangeByScore(key string, min, max interface{}, withscore bool, limit ...interface{}) *Reply {

	var args []interface{}

	args = append(args, key)
	args = append(args, min)
	args = append(args, max)

	if withscore {
		args = append(args, "WITHSCORES")
	}

	if len(limit) == 2 {
		args = append(args, "LIMIT")
		args = append(args, limit[0]) // offset
		args = append(args, limit[1]) // count
	}

	return NewReply(z.conn.Do("ZRevRangeByScore", args...))
}

// ------ 查: 按照 Rank

// ZRange key start stop [WITHSCORES]
// @return 指定排行[start,end]内的成员
func (z *ZsetRds) ZRange(key string, start, stop interface{}, withScore bool) *Reply {
	if withScore {
		return NewReply(z.conn.Do("ZRange", key, start, stop, "WITHSCORES"))
	}
	return NewReply(z.conn.Do("ZRange", key, start, stop))
}

func (z *ZsetRds) ZRevRange(key string, start, stop interface{}, withScore bool) *Reply {
	if withScore {
		return NewReply(z.conn.Do("ZRevRange", key, start, stop, "WITHSCORES"))
	}
	return NewReply(z.conn.Do("ZRevRange", key, start, stop))
}

/* --------- 多个ZSet之间的运算 --------- */

// ZUNIONSTORE
// ZINTERSTORE

/* --------- ZSCAN --------- */

// ZSCAN key cursor [MATCH pattern] [COUNT count]
func (z *ZsetRds) ZScan(key string, cursor int64, pattern string, count int64) *Reply {

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

	return NewReply(z.conn.Do("SScan", args...))
}
