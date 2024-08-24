package util

import "fmt"

func StrToTrueWalletMessage(s string) string {
	hexString := ""
	for _, c := range s {
		hexString += "00" + fmt.Sprintf("%02X", c)
	}
	return hexString
}
