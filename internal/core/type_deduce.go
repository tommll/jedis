package core

import "jedis/internal/constant"

func deduceTypeString(value string) (uint8, uint8) {
	return constant.ObjTypeString, constant.ObjEncodingRaw
}
