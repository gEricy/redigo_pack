package redigo_pack

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHashRds_Hash_Common(t *testing.T) {

	Convey("=====TestHashRds_Hash_Common=====", t, func() {
		Convey("=====TestHashRds_Hash_Common, TEST Result: =====", func() {

			type HSet struct {
				Key   string
				Filed string
				Value int64
			}

			key := "TestHashRds_Hash_Common"

			hs := []HSet{
				{Key: key, Filed: "filed1", Value: 1},
				{Key: key, Filed: "filed2", Value: 2},
				{Key: key, Filed: "filed3", Value: 3},
			}

			// HSet \ HGet
			{
				// HSet
				for _, hash := range hs {
					if err := gConn.Hash.HSet(hash.Key, hash.Filed, hash.Value).error; err != nil {
						panic(err)
					}
				}
				// HGet
				for _, hash := range hs {
					v, err := gConn.Hash.HGet(hash.Key, hash.Filed).Int64()
					if err != nil {
						panic(err)
					}
					So(v == hash.Value, ShouldEqual, true)
				}
			}

			// HMget
			{
				value, err := gConn.Hash.HMget(key, []string{"filed1", "filed2"}).Int64s()
				if err != nil {
					panic(err)
				}
				So(value[0] == 1 && value[1] == 2, ShouldEqual, true)
			}

			// HExists
			{
				exist, err := gConn.Hash.HExists(key, "filed1").Bool()
				if err != nil {
					panic(err)
				}
				So(exist, ShouldEqual, true)
			}

			// HKeys : 获取hash所有字段
			{
				fileds, err := gConn.Hash.HKeys(key).Strings()
				if err != nil {
					panic(err)
				}
				So(len(fileds) == len(hs), ShouldEqual, true)
				for i, v := range hs {
					So(v.Filed == fileds[i], ShouldEqual, true)
				}
			}

			// HLen: 获取hash字段数量
			{
				filedNum, err := gConn.Hash.HLen(key).Int()
				if err != nil {
					panic(err)
				}
				So(filedNum == len(hs), ShouldEqual, true)
			}

			// HVals: 获取hash所有字段值
			{
				values, err := gConn.Hash.HVals(key).Int64s()
				if err != nil {
					panic(err)
				}
				So(len(values) == len(hs), ShouldEqual, true)
				for i, v := range hs {
					So(v.Value == values[i], ShouldEqual, true)
				}
			}

			// HDel
			{
				err := gConn.Hash.HDel(key, []string{"filed1", "filed2"}).error
				if err != nil {
					panic(err)
				}
				exist1, err := gConn.Hash.HExists(key, "filed1").Bool()
				if err != nil {
					panic(err)
				}
				So(exist1, ShouldEqual, false)
				exist2, err := gConn.Hash.HExists(key, "filed2").Bool()
				if err != nil {
					panic(err)
				}
				So(exist2, ShouldEqual, false)
			}

			gConn.Key.Del([]string{key})
		})
	})
}

// map <---> hmset
func TestHashRds_HMsetFromMap(t *testing.T) {

	Convey("=====TestHashRds_HMsetFromMap=====", t, func() {
		Convey("=====TestHashRds_HMsetFromMap, TEST Result: =====", func() {

			key1 := "TestHashRds_HMsetFromMap"
			// hmset: map ---> Hash
			{
				mp := map[interface{}]interface{}{
					"filed1": 1,
					"filed2": 2,
				}

				err := gConn.Hash.HMsetFromMap(key1, mp).error
				if err != nil {
					panic(err)
				}

				for k, v := range mp {
					val, err := gConn.Hash.HGet(key1, k).Int64()
					if err != nil {
						panic(err)
					}
					So(val, ShouldEqual, v)
				}
			}

			gConn.Key.Del([]string{key1})

		})
	})
}

// struct <---> hmset
func TestHashRds_HMsetFromStruct(t *testing.T) {

	Convey("=====TestHashRds_HMsetFromStruct=====", t, func() {
		Convey("=====TestHashRds_HMsetFromStruct, TEST Result: =====", func() {

			type Student struct {
				Name string
				Age  int64
			}

			key := "TestHashRds_HMsetFromStruct"

			// hmset: struct <---> Hash
			{
				var (
					stu = &Student{
						Name: "Tom",
						Age:  24,
					}
				)
				// struct --> HSet <key> { <filed,val>, <filed,val> ... }
				err := gConn.Hash.HMsetFromStruct(key, stu).error
				if err != nil {
					panic(err)
				}
				// HSet <key> { <filed,val>, <filed,val> ... } --> struct
				hvalRes := &Student{}
				err = gConn.Hash.HGetAll(key).ScanStruct(hvalRes)
				if err != nil {
					panic(err)
				}
				So(hvalRes.Name == stu.Name && hvalRes.Age == stu.Age, ShouldEqual, true)
			}

			gConn.Key.Del([]string{key})
		})
	})
}

// User-defined type
func TestHashRds_Hash_UserType(t *testing.T) {
	Convey("=====TestHashRds_Hash_UserType=====", t, func() {
		Convey("=====TestHashRds_Hash_UserType, TEST Result: =====", func() {

			key := "TestHashRds_Hash_UserType"

			type Student struct {
				Name string
				Age  int
			}

			type HSet struct {
				Key   string
				Filed string
				Value []byte // save Student
			}

			// 数据准备
			stu1 := &Student{
				Name: "Stu1",
				Age:  18,
			}
			stu2 := &Student{
				Name: "Stu2",
				Age:  20,
			}
			byteStu1, _ := json.Marshal(stu1)
			byteStu2, _ := json.Marshal(stu2)

			hs := []HSet{
				{Key: key, Filed: "filed1", Value: byteStu1},
				{Key: key, Filed: "filed2", Value: byteStu2},
			}

			// HSet \ HGet
			{
				// HSet
				for _, hash := range hs {
					if err := gConn.Hash.HSet(hash.Key, hash.Filed, hash.Value).error; err != nil {
						panic(err)
					}
				}
				// HGet
				for _, hash := range hs {
					var srcStu Student
					if err := json.Unmarshal(hash.Value, &srcStu); err != nil {
						panic(err)
					}
					v, err := gConn.Hash.HGet(hash.Key, hash.Filed).Bytes()
					if err != nil {
						panic(err)
					}
					var DstStu Student
					if err := json.Unmarshal(v, &DstStu); err != nil {
						panic(err)
					}
					So(srcStu.Name == DstStu.Name && srcStu.Age == DstStu.Age, ShouldEqual, true)
				}
			}

			// HMget
			{
				value, err := gConn.Hash.HMget(key, []string{"filed1", "filed2"}).ByteSlices()
				if err != nil {
					panic(err)
				}
				So(len(value), ShouldEqual, 2)
				for i := range []string{"filed1", "filed2"} {
					var DstStu Student
					if err := json.Unmarshal(value[i], &DstStu); err != nil {
						panic(err)
					}
					var srcStu Student
					if err := json.Unmarshal(hs[i].Value, &srcStu); err != nil {
						panic(err)
					}
					So(srcStu.Name == DstStu.Name && srcStu.Age == DstStu.Age, ShouldEqual, true)
				}
			}

			// HGetall
			{
				reply, err := gConn.Hash.HGetAll(key).StringMap()
				if err != nil {
					panic(err)
				}
				res := make(map[string]Student, 0)
				for filed, value := range reply {
					var stu Student
					if err := json.Unmarshal([]byte(value), &stu); err != nil {
						panic("josn unmarshal fail")
					}
					res[filed] = stu
				}
				// map[filed1:{Stu1 18} filed2:{Stu2 20}]
				t.Log(res)
			}

			gConn.Key.Del([]string{key})
		})
	})
}

func TestHashRds_Hash_HSetnx(t *testing.T) {

	Convey("=====TestHashRds_Hash_HSetnx=====", t, func() {
		Convey("=====TestHashRds_Hash_HSetnx, TEST Result: =====", func() {

			key := "TestHashRds_Hash_HSetnx"

			{
				filed := "filed"
				value := "value"

				reply, err := gConn.Hash.HSetNx(key, filed, value).Result()
				if err != nil {
					panic(err)
				}
				So(reply, ShouldEqual, 1)
				reply, err = gConn.Hash.HSetNx(key, filed, value).Result()
				if err != nil {
					panic(err)
				}
				So(reply, ShouldEqual, 0)
			}

			gConn.Key.Del([]string{key})
		})
	})
}

func TestHashRds_HScanAll_CommType(t *testing.T) {
	Convey("=====TestHashRds_HScanAll_CommType=====", t, func() {
		Convey("=====TestHashRds_HScanAll_CommType, TEST Result: =====", func() {

			key := "TestHashRds_HScanAll_CommType"

			// Hmset: 构造数据
			{
				mm := make(map[interface{}]interface{}, 0)
				for i := 1; i <= 100; i++ {
					mm[i] = i
				}
				if err := gConn.Hash.HMsetFromMap(key, mm).error; err != nil {
					panic(err)
				}
			}

			// HScanAll
			{
				ansMap, err := gConn.Hash.HScanAll(key)
				if err != nil {
					panic(err)
				}

				ans := make(map[int]int, 0)
				for key, val := range ansMap {
					keyInt, _ := strconv.Atoi(key)
					valInt, _ := strconv.Atoi(val)
					ans[keyInt] = valInt
				}
				t.Log("ans = ", ans)
			}

			gConn.Key.Del([]string{key})
		})
	})
}

func TestHashRds_HScanAll_UserType(t *testing.T) {
	Convey("=====TestHashRds_HScanAll_UserType=====", t, func() {
		Convey("=====TestHashRds_HScanAll_UserType, TEST Result: =====", func() {

			key := "TestHashRds_HScanAll_UserType"

			// Hmset
			{
				mm := make(map[interface{}]interface{}, 0)
				for i := int64(1); i <= 10; i++ {
					stu := &Student{
						Name: fmt.Sprintf("name_%d", i),
						Age:  i,
					}
					stuByte, _ := json.Marshal(stu)
					mm[i] = stuByte
				}
				if err := gConn.Hash.HMsetFromMap(key, mm).error; err != nil {
					panic(err)
				}
			}
			// HScanAll
			{
				ansMap, err := gConn.Hash.HScanAll(key)
				if err != nil {
					panic(err)
				}

				ans := make(map[int]Student, 0)
				for key, val := range ansMap {
					keyInt, _ := strconv.Atoi(key)
					var stu Student
					if err := json.Unmarshal([]byte(val), &stu); err != nil {
						panic(err)
					}
					ans[keyInt] = stu
				}
				t.Log("len(ans)=", len(ans))
				for k, v := range ans {
					t.Log(k, v)
				}
			}

			gConn.Key.Del([]string{key})
		})
	})
}

func TestHashRds_Scan(t *testing.T) {
	Convey("=====TestHashRds_Scan=====", t, func() {
		Convey("=====TestHashRds_Scan, TEST Result: =====", func() {

			key := "TestHashRds_Scan"

			// Hmset
			{
				mm := make(map[interface{}]interface{}, 0)
				for i := 1; i <= 10; i++ {
					mm[fmt.Sprintf("key_%d", i)] = i
				}
				if err := gConn.Hash.HMsetFromMap(key, mm).error; err != nil {
					panic(err)
				}
			}
			// HScan
			{
				var (
					cursor  = int64(0)
					pattern = "key_1*"
					count   = int64(1)
				)
				datas, err := gConn.Hash.HScan(key, cursor, pattern, count).Values()
				if err != nil {
					panic(err)
				}
				if len(datas) != 2 {
					panic("HScan return ")
				}
				t.Log("len(datas)=", len(datas))
				// 下一个新的游标
				nextCursor, err := redis.Int64(datas[0], nil)
				if err != nil {
					return
				}
				t.Log("nextCursor=", nextCursor)
				// 结果集
				key2valMap, err := redis.StringMap(datas[1], nil)
				if err != nil {
					panic(err)
				}
				t.Log("len(key2valMap)=", len(key2valMap))
				for filed, value := range key2valMap {
					t.Log("filed:", filed, "   ", "value:", value)
				}
			}

			gConn.Key.Del([]string{key})
		})
	})
}
