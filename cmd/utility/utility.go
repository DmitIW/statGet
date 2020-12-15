package utility

import (
	"log"
	"strconv"
)

func GetUint16(value string) uint16 {
	var (
		converted uint64
		err       error
	)
	if converted, err = strconv.ParseUint(value, 10, 16); err != nil {
		log.Printf("Error on uint converting: %v\n", err)
		return 0
	}
	return uint16(converted)
}
