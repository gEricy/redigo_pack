package redigo_pack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var gConnScript *RedisDB

func init() {
	conn, err := NewRedisDBWithConfig("127.0.0.1", "6379")
	if err != nil {
		panic("conn fail")
	}
	gConnScript = conn
}

func TestStringRds_Script(t *testing.T) {
	Convey("=====TestStringRds_Script=====", t, func() {
		Convey("=====TestStringRds_Script, TEST Result: =====", func() {

			key := "scriptKey"
			gConn.Hash.HMsetFromMap(key, map[interface{}]interface{}{
				1: 1,
				2: 2,
				3: 3,
			})

			script := `
				local val=redis.call('hget', KEYS[1], ARGV[1])
				if val then
					-- 存在
					return 1
				else
					-- 不存在
					redis.call('hset', KEYS[1], ARGV[1], ARGV[1])
					return 0
				end
			`

			ans, err := gConnScript.Script.DoScript(1, script, key, 9).Int64()
			if err != nil {
				panic(err)
			}
			switch ans {
			case 0:
				t.Logf("reply:%d", ans)
			case 1:
				t.Logf("reply:%d", ans)
			}

		})
	})
}

func TestStringRds_Script_Zset(t *testing.T) {
	Convey("=====TestStringRds_Script_Zset=====", t, func() {
		Convey("=====TestStringRds_Script_Zset, TEST Result: =====", func() {

			key := "TestStringRds_Script_Zset"

			score := 100
			member := "uin_1"

			maxCnt := 1

			script := `
				local zsetKeyExists = redis.call('exists', KEYS[1])
				if zsetKeyExists~=1 then
					redis.call('zadd', KEYS[1], ARGV[1], ARGV[2])
					return 1
				end
				local cnt = redis.call('zcard', KEYS[1])
				if (cnt >= tonumber(ARGV[3])) then
					return 2
				end
				redis.call('zadd', KEYS[1], ARGV[1], ARGV[2])
				return 3
			`

			ans, err := gConnScript.Script.DoScript(1, script,
				key,
				score, member,
				maxCnt,
			).Int64()
			if err != nil {
				panic(err)
			}

			t.Logf("reply:%d", ans)
		})
	})
}

func TestStringRds_Script_Zsethmset(t *testing.T) {
	Convey("=====TestStringRds_Script_Zsethmset=====", t, func() {
		Convey("=====TestStringRds_Script_Zsethmset, TEST Result: =====", func() {

			key := "TestStringRds_Script_Zsethmset"

			script := `
				local res=redis.call('hgetall', KEYS[1])
				local ans = {}
				for filed, value in pairs(res) do
    				ans[filed] = value
				end
				return ans
			`

			ans, err := gConnScript.Script.DoScript(1, script,
				key,
			).StringMap()
			if err != nil {
				panic(err)
			}

			t.Logf("结果:%+v", ans)

		})
	})
}
