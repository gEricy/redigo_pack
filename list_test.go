package redigo_pack

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHashRds_List_Common(t *testing.T) {

	Convey("=====TestHashRds_List_Common=====", t, func() {
		Convey("=====TestHashRds_List_Common, TEST Result: =====", func() {

			type Student struct {
				Name string
				Age  int64
			}

			key1 := "TestHashRds_List_Common"
			{
				// LPush
				{
					stuByte, _ := json.Marshal(&Student{
						Name: "LPush",
						Age:  0,
					})
					if err := gConn.List.LPush(key1, []interface{}{1, "2", stuByte}).error; err != nil {
						panic(err)
					}
				}
				// RPush
				{
					stuByte, _ := json.Marshal(&Student{
						Name: "RPush",
						Age:  0,
					})
					if err := gConn.List.RPush(key1, []interface{}{1, "2", stuByte}).error; err != nil {
						panic(err)
					}
				}
				// LPushX
				{
					if err := gConn.List.LPushX(key1, "head").error; err != nil {
						panic(err)
					}
				}
				// RPushX
				{
					if err := gConn.List.RPushX(key1, "tail").error; err != nil {
						panic(err)
					}
				}
				// LInsert
				{
					stuByte, _ := json.Marshal(&Student{
						Name: "LInsert",
						Age:  0,
					})
					gConn.List.LInsert(key1, InsertTypeBefore, "head", stuByte)
					gConn.List.LInsert(key1, InsertTypeAfter, "head", stuByte)
				}
				// Lpop
				{
					gConn.List.LPop(key1)
					gConn.List.RPop(key1)
				}
				// LRem
				{
					gConn.List.LRem(key1, 2, 1)
				}
			}
		})
	})
}

func TestHashRds_List_LPop(t *testing.T) {

	Convey("=====TestHashRds_List_LPop=====", t, func() {
		Convey("=====TestHashRds_List_LPop, TEST Result: =====", func() {

			key1 := "TestHashRds_List_LPop1"
			{
				gConn.List.LPush(key1, []interface{}{1, 2, 3, 4})
				{
					gConn.List.BLPop(1, []interface{}{key1})
					gConn.List.BRPop(1, []interface{}{key1})
				}
			}

			key2 := "TestHashRds_List_LPop2"
			{
				gConn.List.LPush(key2, []interface{}{1, 2, 3, 4})
				{
					gConn.List.BLPop(1, []interface{}{key2, key1})
					gConn.List.BRPop(1, []interface{}{key1, key2})
				}
			}

			gConn.Key.Del([]string{key1, key2})
		})
	})
}
