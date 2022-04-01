package redigo_pack

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type HashRds struct {
	conn redis.Conn
}

/* --------- HSet/HGet --------- */

// HSET <key> <field> <value>
func (h *HashRds) HSet(key string, filed, value interface{}) *Reply {
	return NewReply(h.conn.Do("HSet", key, filed, value))
}

// HGet <key> <filed>
func (h *HashRds) HGet(key string, filed interface{}) *Reply {
	return NewReply(h.conn.Do("HGet", key, filed))
}

/* --------- HSetNx --------- */

// HSetNx <key> <filed> <value>
func (h *HashRds) HSetNx(key string, filed, value interface{}) *Reply {
	return NewReply(h.conn.Do("HSetNx", key, filed, value))
}

/* --------- Hmset : HMSET key field value [field value …]--------- */

// [map] --> hmset
func (h *HashRds) HMsetFromMap(key string, filed2ValueMap map[interface{}]interface{}) *Reply {
	return NewReply(h.conn.Do("hmset", redis.Args{}.Add(key).AddFlat(filed2ValueMap)...))
}

// [struct] --> hmset
func (h *HashRds) HMsetFromStruct(key string, structObj interface{}) *Reply {
	return NewReply(h.conn.Do("hmset", redis.Args{}.Add(key).AddFlat(structObj)...))
}

/* --------- HMget --------- */

// HMGET key field [field …]
func (h *HashRds) HMget(key string, fileds []string) *Reply {
	return NewReply(h.conn.Do("HMget", redis.Args{}.Add(key).AddFlat(fileds)...))
}

// HDel key field [field …]
func (h *HashRds) HDel(key string, fileds []string) *Reply {
	return NewReply(h.conn.Do("HDel", redis.Args{}.Add(key).AddFlat(fileds)...))
}

// HExists: 判断key中是否存在filed
func (h *HashRds) HExists(key string, filed string) *Reply {
	return NewReply(h.conn.Do("HExists", key, filed))
}

// HLen: 返回哈希表key中filed的数量
func (h *HashRds) HLen(key string) *Reply {
	return NewReply(h.conn.Do("HLen", key))
}

func (h *HashRds) HKeys(key string) *Reply {
	return NewReply(h.conn.Do("HKeys", key))
}

func (h *HashRds) HVals(key string) *Reply {
	return NewReply(h.conn.Do("HVals", key))
}

// HIncrBy key filed [addNum] 为指定字段值增加
func (h *HashRds) HIncrBy(key string, filed interface{}, increment interface{}) *Reply {
	return NewReply(h.conn.Do("HIncrBy", key, filed, increment))
}

// HIncrByfloat key filed [addNum] 为指定字段值增加
func (h *HashRds) HIncrByFloat(key string, filed interface{}, increment float64) *Reply {
	return NewReply(h.conn.Do("HIncrByfloat", key, filed, increment))
}

/* --------- HGetAll --------- */

// HGetAll <key>: 获取所有字段及值
func (h *HashRds) HGetAll(key string) *Reply {
	return NewReply(h.conn.Do("HGetall", key))
}

// HScanAll: 使用 HScan 实现 HGetall
func (h *HashRds) HScanAll(key string) (map[string]string, error) {
	// HScan <key> <cursor> COUNT <batchCount>  从游标cursor开始，最大查询batchCount个数据
	scanBatchCbk := func(key string, cursor, batchCount int64) (nextCursor int64, key2valMap map[string]string, err error) {
		datas, err := h.HScan(key, cursor, "", batchCount).Values()
		if err != nil {
			return
		}
		if len(datas) != 2 {
			err = fmt.Errorf("not enough return for HScan, key:%s", key)
			return
		}
		// 下一个新的游标
		nextCursor, err = redis.Int64(datas[0], nil)
		if err != nil {
			return
		}
		// 结果集
		key2valMap, err = redis.StringMap(datas[1], nil)
		if err != nil {
			return
		}
		return
	}

	// 使用 HScan 实现 HGetall
	scanAllCbk := func(key string) (map[string]string, error) {

		batchSize := int64(100)
		nextCursor := int64(0)

		ans := make(map[string]string, 0)

		for {
			nextCursor, key2valMap, err := scanBatchCbk(key, nextCursor, batchSize)
			if err != nil {
				return nil, err
			}

			for key, val := range key2valMap {
				ans[key] = val
			}

			if nextCursor == 0 { // 0 表示遍历完成了~
				break
			}
		}

		return ans, nil
	}

	return scanAllCbk(key)
}

// HScan key cursor [MATCH pattern] [COUNT count]
func (h *HashRds) HScan(key string, cursor int64, pattern string, count int64) *Reply {

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

	return NewReply(h.conn.Do("HScan", args...))
}
