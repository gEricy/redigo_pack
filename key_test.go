package redigo_pack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKeyRds_Key(t *testing.T) {
	type T struct {
		Key   string
		Value float64
	}
	ts := []T{
		{Key: "key1", Value: 1.1},
		{Key: "key2", Value: 2.1},
		{Key: "key3", Value: 2.2},
	}
	for _, one := range ts {
		err := gConn.String.Set(one.Key, one.Value).error
		if err != nil {
			t.Error(err)
		}
	}
	key, err := gConn.Key.RandomKey().String()
	if err != nil {
		t.Error(err)
	}
	for Index, one := range ts {
		if one.Key == key {
			break
		}
		if Index == len(ts) {
			t.Error("randomkey 随机key出错")
		}
	}

	for _, one := range ts {
		err := gConn.Key.Rename(one.Key, one.Key+"1").error
		if err != nil {
			t.Error(err)
		}
	}

	for _, one := range ts {
		exist, err := gConn.Key.Exists(one.Key).Bool()
		if err != nil {
			t.Error(err)
		}
		if exist {
			t.Error("rename 出错")
		}
		exist, err = gConn.Key.Exists(one.Key + "1").Bool()
		if err != nil {
			t.Error(err)
		}
		if !exist {
			t.Error("rename 出错")
		}
	}

	for _, one := range ts {
		err = gConn.Key.Expire(one.Key+"1", 1000).error
		if err != nil {
			t.Error(err)
		}
		ttl, err := gConn.Key.TTL(one.Key + "1").Int64()
		if err != nil {
			t.Error(err)
		}
		if ttl == -1 || ttl == -2 {
			t.Error("过期key失败")
		}
	}

	for _, one := range ts {
		err = gConn.Key.Move(one.Key+"1", 2).error
		if err != nil {
			t.Error(err)
		}
	}

	err = gConn.Db.Select(2).error
	if err != nil {
		t.Error(err)
	}
	for _, one := range ts {
		exist, err := gConn.Key.Exists(one.Key + "1").Bool()
		if err != nil {
			t.Error(err)
		}
		if !exist {
			t.Error("move 出错")
		}
	}
}

func TestHashRds_Keys(t *testing.T) {
	Convey("=====TestHashRds_Keys=====", t, func() {
		Convey("=====TestHashRds_Keys, TEST Result: =====", func() {

			key1 := "TestHashRds_Keys1"
			key2 := "TestHashRds_Keys2"

			{
				gConn.String.Set(key1, 1)
				gConn.String.Set(key2, 2)

				gConn.Key.Del([]string{key1, key2})

				key1Exist, err := gConn.Key.Exists(key1).Bool()
				if err != nil {
					panic(err)
				}
				So(key1Exist, ShouldEqual, false)

				key2Exist, err := gConn.Key.Exists(key2).Bool()
				if err != nil {
					panic(err)
				}
				So(key2Exist, ShouldEqual, false)
			}

			gConn.Key.Del([]string{key1, key2})
		})
	})
}
