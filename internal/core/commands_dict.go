package core

import (
	"errors"
	"fmt"
	"jedis/internal/constant"
	"strconv"
)

func cmdSET(args []string) []byte {
	if len(args) < 2 || len(args) == 3 || len(args) > 4 {
		return Encode(errors.New("(error) ERR wrong number of arguments for 'SET' command"), false)
	}

	var key, value string
	var ttlMs int64 = -1

	fmt.Println(dictStore)

	key, value = args[0], args[1]
	oType, oEnc := deduceTypeString(value)
	if len(args) > 2 {
		ttlSec, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			return Encode(errors.New("(error) ERR value is not an integer or out of range"), false)
		}
		ttlMs = ttlSec * 1000
	}

	dictStore.Set(key, dictStore.NewObj(value, uint64(ttlMs), oType, oEnc))
	return constant.RespOk
}

func cmdGET(args []string) []byte {
	if len(args) != 1 {
		return Encode(errors.New("(error) ERR wrong number of arguments for 'GET' command"), false)
	}

	key := args[0]
	obj := dictStore.Get(key)

	fmt.Println("Key", key, "GET result", obj)

	if obj == nil {
		return constant.RespNil
	}

	if dictStore.HasExpired(obj) {
		return constant.RespNil
	}

	return Encode(obj.Value, false)
}
