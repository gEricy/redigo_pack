package redigo_pack

import (
	"encoding/json"
	"testing"

	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
)

var gConn *RedisDB

type Student struct {
	Name string
	Age  int64
}

func init() {
	conn, err := NewRedisDBWithConfig("127.0.0.1", "6379")
	if err != nil {
		panic("conn fail")
	}
	gConn = conn
}

/* --- 字符串 String --- */

// Get \ Set -- val: 普通类型
func TestStringRds_Set_Get_CommonType(t *testing.T) {

	Convey("=====TestStringRds_Set_Get_CommonType=====", t, func() {
		Convey("=====TestStringRds_Set_Get_CommonType, TEST Result: =====", func() {

			var (
				NotExistKey = "NotExistKey"

				key = "TestStringRds_Set_Get_CommonType"
				val = "defaultValPre"
			)
			// get: 查找不存在的Key
			{
				reply, err := gConn.String.Get(NotExistKey).Result()
				So(err, ShouldEqual, nil)
				So(reply, ShouldEqual, nil) // Get不存在的key: err=nil, reply=nil
			}

			// set\get
			{
				// set <key> <val>
				{
					reply, err := gConn.String.Set(key, val).Result()
					if err != nil {
						panic(err)
					}
					So(reply, ShouldEqual, "OK")
					So(err, ShouldEqual, nil)
				}
				// get <key>
				{
					reply, err := gConn.String.Get(key).String()
					if err != nil && err != redis.ErrNil {
						panic(err)
					}
					So(reply, ShouldEqual, val)
					So(err, ShouldEqual, nil)
				}
			}
			gConn.Key.Del([]string{key})
		})
	})
}

// set\get -- val: 自定义结构体类型
func TestStringRds_Set_Get_UserType(t *testing.T) {

	Convey("=====TestStringRds_Set_Get_UserType=====", t, func() {
		Convey("=====TestStringRds_Set_Get_UserType, TEST Result: =====", func() {

			key := "TestStringRds_Set_Get_CommonType"

			{
				val := &Student{
					Name: "Tom",
					Age:  18,
				}
				// set <key> <val>
				{
					// 1. Marshal序列化
					stuByte, err := json.Marshal(val)
					if err != nil {
						panic(err)
					}
					// 2. set key val
					reply, err := gConn.String.Set(key, stuByte).Result()
					if reply != "OK" || err != nil {
						panic(err)
					}
					So(reply, ShouldEqual, "OK")
					So(err, ShouldEqual, nil)
				}

				// get <key>
				{
					// 1. redis.Bytes()
					valByte, err := gConn.String.Get(key).Bytes()
					if err != nil && err != redis.ErrNil {
						panic(err)
					}
					// 2. UnMarshal
					if err != redis.ErrNil {
						var valStu Student
						if err := json.Unmarshal(valByte, &valStu); err != nil {
							panic(err)
						}
						So(valStu.Name, ShouldEqual, "Tom")
						So(valStu.Age, ShouldEqual, 18)
					}
				}
			}
			gConn.Key.Del([]string{key})
		})
	})
}

// mset \ mget --- val: 普通类型
func TestStringRds_MSet_MGet_CommonType(t *testing.T) {

	Convey("=====TestStringRds_MSet_MGet_CommonType=====", t, func() {
		Convey("=====TestStringRds_MSet_MGet_CommonType, TEST Result: =====", func() {

			const (
				// key
				key1        = "TestStringRds_MSet_MGet_CommonType1"
				key2        = "TestStringRds_MSet_MGet_CommonType2"
				NotExistKey = "NotExistKey"
				// val
				val1 = "val1"
				val2 = "val2"
			)
			// 1. val类型是字符串 -- mset\mget
			{
				key2valMap := map[string]interface{}{
					key1: val1,
					key2: val2,
				}
				reply, err := gConn.String.Mset(key2valMap).Result()
				if reply != "OK" || err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(key2valMap), "输出:", reply, err)
					panic("mset fail")
				}
			}
			// mget <key1> <key2> ...
			{
				keys := []string{
					key1,
					key2,
					NotExistKey,
				}
				reply, err := gConn.String.Mget(keys).Strings()
				if err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(keys), "输出:", reply, err)
					panic("mget fail")
				}
				So(err, ShouldEqual, nil)
				So(len(reply), ShouldEqual, 3)
				So(reply[0], ShouldEqual, val1)
				So(reply[1], ShouldEqual, val2)
				So(reply[2], ShouldEqual, "")
			}

			gConn.Key.Del([]string{key1, key2, NotExistKey})
		})
	})
}

// mset/mget -- val: 自定义结构体类型
func TestStringRds_MSet_MGet_UserType(t *testing.T) {

	Convey("=====TestStringRds_MSet_MGet_UserType=====", t, func() {
		Convey("=====TestStringRds_MSet_MGet_UserType, TEST Result: =====", func() {

			const (
				key1        = "TestStringRds_MSet_MGet_UserType1"
				key2        = "TestStringRds_MSet_MGet_UserType2"
				NotExistKey = "NotExistKey"
			)
			// 1. val类型是字符串 -- mset\mget
			const (
				valName1 = "name_1"
				valAge1  = 18
				valName2 = "name_2"
				valAge2  = 19
			)
			{
				byteVal1, _ := json.Marshal(&Student{
					Name: valName1,
					Age:  valAge1,
				})
				byteVal2, _ := json.Marshal(&Student{
					Name: valName2,
					Age:  valAge2,
				})
				key2valMap := map[string]interface{}{
					key1: byteVal1,
					key2: byteVal2,
				}
				reply, err := gConn.String.Mset(key2valMap).Result()
				if reply != "OK" || err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(key2valMap), "输出:", reply, err)
					panic("mset fail")
				}
			}
			// mget <key1> <key2> ...
			{
				keys := []string{
					key1,
					key2,
					NotExistKey, // 不存在的Key
				}
				reply, err := gConn.String.Mget(keys).ByteSlices()
				if err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(keys), "输出:", reply, err)
					panic("mget fail")
				}
				So(err, ShouldEqual, nil)
				So(len(reply), ShouldEqual, 3)

				// key1
				var stu1 Student
				if err := json.Unmarshal(reply[0], &stu1); err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(keys), "输出:", reply, err)
					panic(err)
				}
				So(stu1.Name, ShouldEqual, valName1)
				So(stu1.Age, ShouldEqual, valAge1)
				// key2
				var stu2 Student
				if err := json.Unmarshal(reply[1], &stu2); err != nil {
					panic(err)
				}
				So(stu2.Name, ShouldEqual, valName2)
				So(stu2.Age, ShouldEqual, valAge2)
				// key3
				So(reply[2], ShouldEqual, nil) // 不存在的key，返回的val=nil
			}

			gConn.Key.Del([]string{key1, key2, NotExistKey})
		})
	})
}

// MSetNX -- val: 自定义结构体类型
func TestStringRds_MSetNX_UserType(t *testing.T) {

	Convey("=====TestStringRds_MSetNX_UserType=====", t, func() {
		Convey("=====TestStringRds_MSetNX_UserType, TEST Result: =====", func() {

			const (
				key1        = "TestStringRds_MSetNX_UserType1"
				key2        = "TestStringRds_MSetNX_UserType2"
				NotExistKey = "NotExistKey"
			)
			// 1. val类型是字符串 -- mset\mget
			const (
				valName1 = "name_1"
				valAge1  = 18
				valName2 = "name_2"
				valAge2  = 19
			)
			{
				byteVal1, _ := json.Marshal(&Student{
					Name: valName1,
					Age:  valAge1,
				})
				byteVal2, _ := json.Marshal(&Student{
					Name: valName2,
					Age:  valAge2,
				})
				key2valMap := map[string]interface{}{
					key1: byteVal1,
					key2: byteVal2,
				}
				reply, err := gConn.String.MsetNx(key2valMap).Int()
				if err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(key2valMap), "输出:", reply, err)
					panic("mset fail")
				}
				So(reply, ShouldEqual, 1) // 1: 所有key设置都成功

				reply, err = gConn.String.MsetNx(map[string]interface{}{key1: byteVal1, NotExistKey: NotExistKey}).Int()
				if err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(key2valMap), "输出:", reply, err)
					panic("mset fail")
				}
				So(reply, ShouldEqual, 0) // 0: 存在一个key设置失败 (key1已经存在，设置失败)
			}
			// mget <key1> <key2> ...
			{
				keys := []string{
					key1,
					key2,
					NotExistKey, // 不存在的Key
				}
				reply, err := gConn.String.Mget(keys).ByteSlices()
				if err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(keys), "输出:", reply, err)
					panic("mget fail")
				}
				So(err, ShouldEqual, nil)
				So(len(reply), ShouldEqual, 3)

				// key1
				var stu1 Student
				if err := json.Unmarshal(reply[0], &stu1); err != nil {
					t.Log("输入:", redis.Args{}.AddFlat(keys), "输出:", reply, err)
					panic(err)
				}
				So(stu1.Name, ShouldEqual, valName1)
				So(stu1.Age, ShouldEqual, valAge1)
				// key2
				var stu2 Student
				if err := json.Unmarshal(reply[1], &stu2); err != nil {
					panic(err)
				}
				So(stu2.Name, ShouldEqual, valName2)
				So(stu2.Age, ShouldEqual, valAge2)
				// key3
				So(reply[2], ShouldEqual, nil) // 不存在的key，返回的val=nil
			}

			gConn.Key.Del([]string{key1, key2, NotExistKey})
		})
	})
}

func TestStringRds_Set_Nx_Ex(t *testing.T) {
	// 1. SetNx

	{
		Convey("=====TestStringRds_Set_Nx_Ex=====", t, func() {
			Convey("=====TestStringRds_Set_Nx_Ex, TEST Result: =====", func() {

				key := "TestStringRds_Set_Nx_Ex"
				val := "val"

				// 首次，1: 成功
				{
					reply, err := gConn.String.SetNx(key, val).Int()
					So(err, ShouldEqual, nil)
					So(reply, ShouldEqual, 1)
				}
				// 再次，0: 失败
				{
					reply, err := gConn.String.SetNx(key, val).Int()
					So(err, ShouldEqual, nil)
					So(reply, ShouldEqual, 0)
				}

				gConn.Key.Del([]string{key})
			})
		})
	}
}
