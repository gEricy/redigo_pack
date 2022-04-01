package redigo_pack

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

type ListRds struct {
	conn redis.Conn
}

/* --------- 增 --------- */

// --- LPush \ RPush

// LPush key value [value …]
func (l *ListRds) LPush(key string, values []interface{}) *Reply {

	var args []interface{}

	args = append(args, key)
	args = append(args, values...)

	return NewReply(l.conn.Do("LPush", args...))
}

// RPush key value [value …]
func (l *ListRds) RPush(key string, values []interface{}) *Reply {

	var args []interface{}

	args = append(args, key)
	args = append(args, values...)

	return NewReply(l.conn.Do("RPush", args...))
}

// --- LPushX \ RPushX

// LPushX key value : 将值 value 插入到列表 key 的表头，当且仅当 key 存在并且是一个列表
func (l *ListRds) LPushX(key string, value interface{}) *Reply {
	return NewReply(l.conn.Do("LPushX", key, value))
}

// RPushX key value : 将值 value 插入到列表 key 的表尾，当且仅当 key 存在并且是一个列表
func (l *ListRds) RPushX(key string, value interface{}) *Reply {
	return NewReply(l.conn.Do("RPushX", key, value))
}

// --- LInsert

// LInsert key BEFORE|AFTER pivot value : 将值 value 插入到列表 key 当中，位于 {值pivot} 之前或之后

type LInsertType string

const (
	InsertTypeBefore LInsertType = "BEFORE"
	InsertTypeAfter  LInsertType = "AFTER"
)

func (l *ListRds) LInsert(key string, InsertType LInsertType, pivot, value interface{}) *Reply {
	switch InsertType {
	case InsertTypeBefore, InsertTypeAfter:
	default:
		return NewReply(nil, errors.New("LInsertType is invaild"))
	}
	return NewReply(l.conn.Do("LInsert", key, InsertType, pivot, value))
}

/* --------- Pop 删 --------- */

// LPop key
func (l *ListRds) LPop(key string) *Reply {
	return NewReply(l.conn.Do("LPop", key))
}

// RPop key
func (l *ListRds) RPop(key string) *Reply {
	return NewReply(l.conn.Do("RPop", key))
}

// ---  BLPop \ BRPop - 阻塞: 当给定多个 key 参数时，按参数 key 的先后顺序依次检查各个列表，弹出第一个非空列表的头元素

// BLPop key [key …] timeout
func (l *ListRds) BLPop(timeout int64, keys []interface{}) *Reply {

	var args []interface{}

	args = append(args, keys...)
	args = append(args, timeout)

	return NewReply(l.conn.Do("BLPop", args...))
}

// BRPop key [key …] timeout
func (l *ListRds) BRPop(timeout int64, keys []interface{}) *Reply {

	var args []interface{}

	args = append(args, keys...)
	args = append(args, timeout)

	return NewReply(l.conn.Do("BRPop", args...))
}

// ---  LRem

// LRem key count value: 根据参数 count 的值，移除列表中与参数 value 相等的元素
//    1. count > 0 : 从表头开始向表尾搜索，移除与 value 相等的元素，数量为 count
//    2. count < 0 : 从表尾开始向表头搜索，移除与 value 相等的元素，数量为 count 的绝对值
//    3. count = 0 : 移除表中所有与 value 相等的值
func (l *ListRds) LRem(key string, count int64, value interface{}) *Reply {
	return NewReply(l.conn.Do("LRem", key, count, value))
}

// ---  LTrim

// LTrim key start stop: 修剪 - 让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除
func (l *ListRds) LTrim(key string, start, stop int64) *Reply {
	return NewReply(l.conn.Do("LTrim", key, start, stop))
}

/* --------- SET 改 --------- */

// LSet key index value : 将列表 key 下标为 index 的元素的值设置为 value
func (l *ListRds) LSet(key string, index int64, value interface{}) *Reply {
	return NewReply(l.conn.Do("LSet", key, index, value))
}

/* --------- LIndex \ RANGE 查 --------- */

// LLen key
func (l *ListRds) LLen(key string) *Reply {
	return NewReply(l.conn.Do("LLen", key))
}

// LIndex key index : 返回列表 key 中，下标为 index 的元素
func (l *ListRds) LIndex(key string, index int64) *Reply {
	return NewReply(l.conn.Do("LIndex", key, index))
}

// LRange key start stop : 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定
func (l *ListRds) LRange(key string, start, stop int64) *Reply {
	return NewReply(l.conn.Do("LRange", key, start, stop))
}

/* --------- 原子操作: pop + push --------- */

// RpopLpush source destination : 弹出source尾元素，头插入到destination，返回“被弹出的元素”
func (l *ListRds) RpopLpush(key, source, destination string) *Reply {
	return NewReply(l.conn.Do("RpopLpush", key, source, destination))
}

// BRpopLpush source destination timeout
func (l *ListRds) BRpopLpush(key, source, destination string, timeout int64) *Reply {
	return NewReply(l.conn.Do("BRpopLpush", key, source, destination, timeout))
}
