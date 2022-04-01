package redigo_pack

import (
	"fmt"
	"testing"
)

func TestBitRds_Bitop(t *testing.T) {

	cli := gConn.Bit

	var (
		bitKey1 = "bits-1"
		bitKey2 = "bits-2"
		dstKey  = "dstKey"
	)

	// bitKey1
	{
		cli.SetBit(bitKey1, 0, 1)
		cli.SetBit(bitKey1, 3, 1)

		fmt.Println(bitKey1)
		var bit1Vec []int64
		for i := int64(3); i >= 0; i-- {
			bit, err := cli.GetBit(bitKey1, i).Int64()
			if err != nil {
				panic(err)
			}
			bit1Vec = append(bit1Vec, bit)
		}
		t.Log(bit1Vec)
	}

	// bitKey2
	{
		cli.SetBit(bitKey2, 0, 1)
		cli.SetBit(bitKey2, 1, 1)
		cli.SetBit(bitKey2, 3, 1)

		fmt.Println(bitKey2)
		var bit2Vec []int64
		for i := int64(3); i >= 0; i-- {
			bit, err := cli.GetBit(bitKey2, i).Int64()
			if err != nil {
				panic(err)
			}
			bit2Vec = append(bit2Vec, bit)
		}
		t.Log(bit2Vec)
	}

	// dstKey
	{
		// And
		{
			_, err := cli.Bitop(BitopTypeAnd, dstKey, bitKey1, bitKey2).Int64()
			if err != nil {
				panic(err)
			}

			fmt.Println(dstKey)
			var dstKeyVec []int64
			for i := int64(3); i >= 0; i-- {
				bit, err := cli.GetBit(dstKey, i).Int64()
				if err != nil {
					panic(err)
				}
				dstKeyVec = append(dstKeyVec, bit)
			}
			t.Log("And:", dstKeyVec)
		}
		// Or
		{
			dstKey := "dstKey"
			_, err := cli.Bitop(BitopTypeOr, dstKey, bitKey1, bitKey2).Int64()
			if err != nil {
				panic(err)
			}

			fmt.Println(dstKey)
			var dstKeyVec []int64
			for i := int64(3); i >= 0; i-- {
				bit, err := cli.GetBit(dstKey, i).Int64()
				if err != nil {
					panic(err)
				}
				dstKeyVec = append(dstKeyVec, bit)
			}
			t.Log("Or:", dstKeyVec)
		}
		// Xor
		{
			dstKey := "dstKey"
			_, err := cli.Bitop(BitopTypeXor, dstKey, bitKey1, bitKey2).Int64()
			if err != nil {
				panic(err)
			}

			fmt.Println(dstKey)
			var dstKeyVec []int64
			for i := int64(3); i >= 0; i-- {
				bit, err := cli.GetBit(dstKey, i).Int64()
				if err != nil {
					panic(err)
				}
				dstKeyVec = append(dstKeyVec, bit)
			}
			t.Log("Xor:", dstKeyVec)
		}
		// Not
		{
			{
				dstKey := "dstKey"
				_, err := cli.Bitop(BitopTypeNot, dstKey, bitKey1).Int64()
				if err != nil {
					panic(err)
				}

				fmt.Println(dstKey)
				var dstKeyVec []int64
				for i := int64(3); i >= 0; i-- {
					bit, err := cli.GetBit(dstKey, i).Int64()
					if err != nil {
						panic(err)
					}
					dstKeyVec = append(dstKeyVec, bit)
				}
				t.Log("Not bitKey1:", dstKeyVec)
			}
			{
				dstKey := "dstKey"
				_, err := cli.Bitop(BitopTypeNot, dstKey, bitKey2).Int64()
				if err != nil {
					panic(err)
				}

				fmt.Println(dstKey)
				var dstKeyVec []int64
				for i := int64(3); i >= 0; i-- {
					bit, err := cli.GetBit(dstKey, i).Int64()
					if err != nil {
						panic(err)
					}
					dstKeyVec = append(dstKeyVec, bit)
				}
				t.Log("Not bitKey2:", dstKeyVec)
			}
		}
	}

	{
		gConn.Key.Del([]string{bitKey1, bitKey2, dstKey})
	}
}
