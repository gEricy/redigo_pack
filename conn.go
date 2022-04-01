package redigo_pack

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	String StringRds
	List   ListRds
	Hash   HashRds
	Key    KeyRds
	Set    SetRds
	Zset   ZsetRds
	Bit    BitRds
	Db     DbRds
}

func NewRedisDBWithConfig(ip, port string) (*RedisDB, error) {

	// 连接
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return nil, err
	}

	// 创建结构体
	redisDB := &RedisDB{
		String: StringRds{
			conn: conn,
		},
		List: ListRds{
			conn: conn,
		},
		Hash: HashRds{
			conn: conn,
		},
		Key: KeyRds{
			conn: conn,
		},
		Set: SetRds{
			conn: conn,
		},
		Zset: ZsetRds{
			conn: conn,
		},
		Bit: BitRds{
			conn: conn,
		},
		Db: DbRds{
			conn: conn,
		},
	}

	return redisDB, nil
}
