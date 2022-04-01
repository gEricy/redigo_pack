package redigo_pack

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// ZAdd \ ZRem \ ZScore
func Test_ZSetRds_ZAdd(t *testing.T) {
	Convey("=====Test_ZSetRds_ZAdd=====", t, func() {
		Convey("=====Test_ZSetRds_ZAdd, TEST Result: =====", func() {

			cli := gConn.Zset

			key := "zset_key"
			{
				// ZAddFromMap
				{
					cli.ZAddFromMap(key, map[interface{}]interface{}{
						"uin1": 1,
						"uin2": 2,
					})
					uin_1_score, _ := cli.ZScore(key, "uin1").Int()
					So(uin_1_score, ShouldEqual, 1)
					uin_2_score, _ := cli.ZScore(key, "uin2").Int()
					So(uin_2_score, ShouldEqual, 2)
					cli.ZRem(key, []string{"uin1", "uin_2"})
					uin_1_score1, _ := cli.ZScore(key, "uin1").Result()
					So(uin_1_score1, ShouldEqual, nil)
					uin_2_score, _ = cli.ZScore(key, "uin2").Int()
					So(uin_2_score, ShouldEqual, 2)
					cli.ZAddFromMap(key, map[interface{}]interface{}{
						"uin1": 1,
						"uin2": 2,
						"uin3": 3,
					})
					cli.ZRem(key, []string{"uin1", "uin2"})
					uin_1_score1, _ = cli.ZScore(key, "uin1").Result()
					So(uin_1_score1, ShouldEqual, nil)
					uin_2_score1, _ := cli.ZScore(key, "uin2").Result()
					So(uin_2_score1, ShouldEqual, nil)
				}
			}
			gConn.Key.Del([]string{key})
		})
	})
}

func Test_ZSetRds_ZRem(t *testing.T) {
	Convey("=====Test_ZSetRds_ZRem=====", t, func() {
		Convey("=====Test_ZSetRds_ZRem, TEST Result: =====", func() {

			cli := gConn.Zset

			key := "zset_key"
			{
				cli.ZAddFromMap(key, map[interface{}]interface{}{
					"uin_1": 1,
					"uin_2": 2,
					"uin_3": 3,
					"uin_4": 4,
				})

				{
					cli.ZRemRangeByRank(key, 1, 2)

					len, _ := cli.ZCard(key).Int()
					So(len, ShouldEqual, 2)
				}
				cli.ZAddFromMap(key, map[interface{}]interface{}{
					"uin_1": 1,
					"uin_2": 2,
					"uin_3": 3,
					"uin_4": 4,
				})
				cli.ZRemRangeByScore(key, 1, 3)
			}
			{
				cli.ZAddFromMap(key, map[interface{}]interface{}{
					"uin_1": 1,
					"uin_2": 2,
					"uin_3": 3,
					"uin_4": 4,
				})
				cli.ZIncrBy(key, 10, "uin_1")
				cli.ZIncrBy(key, 10, "uin_5")
				rank, _ := cli.ZRank(key, "uin_2").Int()
				So(rank, ShouldEqual, 0)
				revrank, _ := cli.ZRevRank(key, "uin_3").Int()
				So(revrank, ShouldEqual, 3)

				cnt, _ := cli.ZCount(key, 2, 6).Int()
				So(cnt, ShouldEqual, 3)

				{
					mems, _ := cli.ZRangeByScore(key, 2, 6, false).Strings()
					So(mems[0], ShouldEqual, "uin_2")
					So(mems[1], ShouldEqual, "uin_3")
					So(mems[2], ShouldEqual, "uin_4")
				}
				{
					err := cli.ZRangeByScore(key, 2, 6, true).error
					So(err, ShouldEqual, nil)
				}
				{
					err := cli.ZRangeByScore(key, 2, 6, true, 0, 100).error
					So(err, ShouldEqual, nil)
				}
				{
					err := cli.ZRange(key, 2, 6, false).error
					So(err, ShouldEqual, nil)
					err = cli.ZRange(key, 2, 6, true).error
					So(err, ShouldEqual, nil)
				}
			}
			gConn.Key.Del([]string{key})
		})
	})
}
