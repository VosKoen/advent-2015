package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	secret := "yzbqklnj"
	foundFive := false
	foundSix := false
	count := 0
	for !foundFive || !foundSix {
		count += 1
		try := secret + strconv.FormatInt(int64(count), 10)
		md5Hash := md5.Sum([]byte(try))
		stringRepresentation := hex.EncodeToString(md5Hash[:])
		if !foundFive && stringRepresentation[:5] == "00000" {
			fmt.Println("Five zeros:", stringRepresentation)
			fmt.Println(count, try)
			foundFive = true
		}
		if !foundSix && stringRepresentation[:6] == "000000" {
			fmt.Println("Siz zeros:", stringRepresentation)
			fmt.Println(count, try)
			foundSix = true
		}
	}
}
