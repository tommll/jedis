package config

var Host = "0.0.0.0"
var Port = 8080
var MaxConnections = 20000
var KeyNumberLimit = 5000000
var DefaultTTLSeconds = 3600

const (
	EvictFirst int = 0
	LRU            = 1
	LFU            = 2
)

var EvictStrategy = EvictFirst
