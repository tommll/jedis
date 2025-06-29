package constant

var RespNil = []byte("$-1\r\n")
var RespOk = []byte("+OK\r\n")
var RespZero = []byte(":0\r\n")
var RespOne = []byte(":1\r\n")
var RespEmptyArray = []byte("*0\r\n")
var TtlKeyNotExist = []byte(":-2\r\n")
var TtlKeyExistNoExpire = []byte(":-1\r\n")

const (
	StatusWaiting      = 0
	StatusBusy         = 1
	StatusShuttingDown = 2
)

const (
	ObjTypeString  uint8 = 0
	ObjTypeSet     uint8 = 1
	ObjTypeZSet    uint8 = 2
	ObjTypeGeoHash uint8 = 3
)

const ObjEncodingRaw uint8 = 0
const ObjEncodingInt uint8 = 1

const CRLF = "\r\n"
