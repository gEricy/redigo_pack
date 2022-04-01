package redigo_pack

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	conn, err := NewRedisDBWithConfig("127.0.0.1", "6379")
	if err != nil {
		panic("conn fail")
	}
	gConn = conn
}

func TestStringRds_Set_Common(t *testing.T) {
	Convey("=====TestStringRds_Set_Common=====", t, func() {
		Convey("=====TestStringRds_Set_Common, TEST Result: =====", func() {

			key := "TestStringRds_Set_Common"
			{
				addList := []interface{}{1, 2, 3, 4}
				_, err := gConn.Set.SAdd(key, addList).Result()
				if err != nil {
					panic(err)
				}
				members, err := gConn.Set.SMembers(key).Int64s()
				if err != nil {
					panic(err)
				}
				So(len(members), ShouldEqual, len(addList))
				// SCard
				lenSet, err := gConn.Set.SCard(key).Int64()
				if err != nil {
					panic(err)
				}
				So(lenSet, ShouldEqual, len(addList))
				// SIsMember
				isExists, err := gConn.Set.SIsMember(key, addList[0]).Bool()
				if err != nil {
					panic(err)
				}
				So(isExists, ShouldEqual, true)
				// SRem
				if err := gConn.Set.SRem(key, addList).error; err != nil {
					panic(err)
				}
				// SCard
				lenSet, err = gConn.Set.SCard(key).Int64()
				if err != nil {
					panic(err)
				}
				So(lenSet, ShouldEqual, 0)
			}
			gConn.Key.Del([]string{key})
		})
	})
}

func sliceIsSame(a, b []int64) bool {
	lenA := len(a)
	lenB := len(b)
	if lenA != lenB {
		return false
	}
	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestStringRds_Set_Multi(t *testing.T) {
	Convey("=====TestStringRds_Set_Multi=====", t, func() {
		Convey("=====TestStringRds_Set_Multi, TEST Result: =====", func() {

			key1 := "TestStringRds_Set_Multi1"
			key2 := "TestStringRds_Set_Multi2"
			// init data
			{
				// set1
				addList1 := []interface{}{1, 2, 3, 4}
				gConn.Set.SAdd(key1, addList1)
				// set2
				addList2 := []interface{}{1, 2, 5}
				gConn.Set.SAdd(key2, addList2)
			}
			// SDiff
			{
				diff, err := gConn.Set.SDiff([]string{key1, key2}).Int64s()
				if err != nil {
					panic(err)
				}
				So(sliceIsSame(diff, []int64{3, 4}), ShouldEqual, true)
			}
			// SInter
			{
				inter, err := gConn.Set.SInter([]string{key1, key2}).Int64s()
				if err != nil {
					panic(err)
				}
				So(sliceIsSame(inter, []int64{1, 2}), ShouldEqual, true)
			}
			// SUnion
			{
				union, err := gConn.Set.SUnion([]string{key1, key2}).Int64s()
				if err != nil {
					panic(err)
				}
				So(sliceIsSame(union, []int64{1, 2, 3, 4, 5}), ShouldEqual, true)
			}
			gConn.Key.Del([]string{key1, key2})
		})
	})
}

func TestStringRds_Set_SScan(t *testing.T) {
	Convey("=====TestStringRds_Set_SScan=====", t, func() {
		Convey("=====TestStringRds_Set_SScan, TEST Result: =====", func() {

			key := "TestStringRds_Set_SScan"
			{
				addList := []interface{}{1, 2, 3, 4}
				gConn.Set.SAdd(key, addList)
			}
			{
				var (
					cursor  = int64(0)
					pattern = "*"
					count   = int64(1)
				)
				datas, err := gConn.Set.SScan(key, cursor, pattern, count).Values()
				if err != nil {
					panic(err)
				}
				if len(datas) != 2 {
					panic("SScan fail, len(datas)!=2")
				}
				t.Log("len(datas)=", len(datas))
				// 下一个新的游标
				nextCursor, err := redis.Int64(datas[0], nil)
				if err != nil {
					return
				}
				t.Log("nextCursor=", nextCursor)
				// 结果集
				ans, err := redis.Int64s(datas[1], nil)
				if err != nil {
					panic(err)
				}
				t.Log("len(ans)=", len(ans))
				for i, elem := range ans {
					t.Log("i:", i, "   ", "elem:", elem)
				}
			}
			gConn.Key.Del([]string{key})
		})
	})
}
