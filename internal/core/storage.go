package core

import "jedis/internal/data_structures"

var dictStore *data_structures.Dict

func init() {
	dictStore = data_structures.NewDict()
}
