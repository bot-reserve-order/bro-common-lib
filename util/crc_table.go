package util

import "github.com/sigurn/crc16"

var crc16Table = crc16.MakeTable(crc16.CRC16_CCITT_FALSE)

func Crc16Hash() crc16.Hash16 {
	return crc16.New(crc16Table)
}