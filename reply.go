package redigo_pack

import (
	"github.com/gomodule/redigo/redis"
)

type Reply struct {
	reply interface{}
	error error
}

func NewReply(rp interface{}, err error) *Reply {
	return &Reply{reply: rp, error: err}
}

func (r *Reply) Result() (interface{}, error) {
	return r.reply, r.error
}

func (r *Reply) Error() error {
	return r.error
}

/* ------------ 单个 ------------*/

func (r *Reply) Int() (int, error) {
	return redis.Int(r.reply, r.error)
}
func (r *Reply) Int64() (int64, error) {
	return redis.Int64(r.reply, r.error)
}

func (r *Reply) Uint64() (uint64, error) {
	return redis.Uint64(r.reply, r.error)
}

func (r *Reply) Float64() (float64, error) {
	return redis.Float64(r.reply, r.error)
}

func (r *Reply) String() (string, error) {
	return redis.String(r.reply, r.error)
}

func (r *Reply) Bytes() ([]byte, error) {
	return redis.Bytes(r.reply, r.error)
}

func (r *Reply) Bool() (bool, error) {
	return redis.Bool(r.reply, r.error)
}

/* ------------ Positions \ SlowLogs ------------*/

func (r *Reply) Positions() ([]*[2]float64, error) {
	return redis.Positions(r.reply, r.error)
}

func (r *Reply) SlowLogs() ([]redis.SlowLog, error) {
	return redis.SlowLogs(r.reply, r.error)
}

/* ------------ Slice ------------*/

func (r *Reply) Float64s() ([]float64, error) {
	return redis.Float64s(r.reply, r.error)
}

func (r *Reply) Strings() ([]string, error) {
	return redis.Strings(r.reply, r.error)
}

func (r *Reply) ByteSlices() ([][]byte, error) {
	return redis.ByteSlices(r.reply, r.error)
}

func (r *Reply) Int64s() ([]int64, error) {
	return redis.Int64s(r.reply, r.error)
}

func (r *Reply) Uint64s() ([]uint64, error) {
	return redis.Uint64s(r.reply, r.error)
}

func (r *Reply) Ints() ([]int, error) {
	return redis.Ints(r.reply, r.error)
}

/* ------------ Map ------------*/

func (r *Reply) StringMap() (map[string]string, error) {
	return redis.StringMap(r.reply, r.error)
}

func (r *Reply) IntMap() (map[string]int, error) {
	return redis.IntMap(r.reply, r.error)
}

func (r *Reply) Int64Map() (map[string]int64, error) {
	return redis.Int64Map(r.reply, r.error)
}

func (r *Reply) Uint64Map() (map[string]uint64, error) {
	return redis.Uint64Map(r.reply, r.error)
}

/* ------------ Map ------------*/

// Values is a helper that converts an array command reply to a []interface{}.
// If err is not equal to nil, then Values returns nil, err. Otherwise, Values
// converts the reply as follows:
//
//  Reply type      Result
//  array           reply, nil
//  nil             nil, ErrNil
//  other           nil, error
func (r *Reply) Values() ([]interface{}, error) {
	return redis.Values(r.reply, r.error)
}

// obj为一个指针对象
// HGetALL
func (r *Reply) ScanStruct(obj interface{}) error {
	v, err := r.Values()
	if err != nil {
		return err
	}
	return redis.ScanStruct(v, obj)
}
