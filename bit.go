package redigo_pack

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

/* --- BitMap --- */

type BitRds struct {
	conn redis.Conn
}

// SetBit key offset value
func (b *BitRds) SetBit(key string, offset, value int64) *Reply {
	return NewReply(b.conn.Do("SetBit", key, offset, value))
}

// GetBit key offset
func (b *BitRds) GetBit(key string, offset int64) *Reply {
	return NewReply(b.conn.Do("GetBit", key, offset))
}

// BITCOUNT key [start] [end]
func (b *BitRds) BitCount(key string, interval ...int64) *Reply {
	if len(interval) == 2 {
		return NewReply(b.conn.Do("BitCount", key, interval[0], interval[1]))
	}
	return NewReply(b.conn.Do("BitCount", key))
}

/* --------- BITOP --------- */

// BITOP operation destkey key [key …]
//       operation = { and、or、xor、not }
// keys多个: and、or、xor
// keys单个: not

type BitopType string

const (
	BitopTypeAnd BitopType = "And"
	BitopTypeOr  BitopType = "Or"
	BitopTypeXor BitopType = "Xor"
	BitopTypeNot BitopType = "Not"
)

func (b *BitRds) Bitop(opt BitopType, destKey string, keys ...interface{}) *Reply {

	switch opt {
	case
		BitopTypeAnd,
		BitopTypeOr,
		BitopTypeXor,
		BitopTypeNot:
	default:
		return NewReply(nil, errors.New("bitop operator type is invaild"))
	}

	var args []interface{}

	args = append(args, opt)
	args = append(args, destKey)
	args = append(args, keys...)

	return NewReply(b.conn.Do("Bitop", args...))
}
