package data_structures

import (
	"jedis/internal/config"
	"time"
)

type Obj struct {
	Value        interface{}
	TypeEncoding uint8
}

type Dict struct {
	items        map[string]*Obj
	expiredItems map[*Obj]uint64
}

func NewDict() *Dict {
	return &Dict{
		items:        make(map[string]*Obj),
		expiredItems: make(map[*Obj]uint64),
	}
}

func (d *Dict) NewObj(value interface{}, ttlMs uint64, oType uint8, oEnc uint8) *Obj {
	obj := &Obj{
		Value:        value,
		TypeEncoding: oType,
	}

	if ttlMs > 0 {
		d.SetExpiredTime(obj, ttlMs)
	}

	return obj
}

func (d *Dict) GetExpiredTime(obj *Obj) (uint64, bool) {
	exp, exist := d.expiredItems[obj]
	if !exist {
		return 0, exist
	}
	return exp, exist
}

func (d *Dict) SetExpiredTime(obj *Obj, ttlMs uint64) {
	d.expiredItems[obj] = uint64(time.Now().UnixMilli()) + ttlMs
}

func (d *Dict) HasExpired(obj *Obj) bool {
	exp, exist := d.expiredItems[obj]
	if !exist {
		return false
	}
	return exp <= uint64(time.Now().UnixMilli())
}

func (d *Dict) Del(key string) bool {
	if obj, exist := d.items[key]; exist {
		delete(d.items, key)
		delete(d.expiredItems, obj)
		return true
	}

	return false
}

func (d *Dict) Set(key string, value *Obj) {
	if len(d.items) >= config.KeyNumberLimit {
		d.evict()
	}
	d.items[key] = value
}

func (d *Dict) Get(key string) *Obj {
	v := d.items[key]

	if v != nil {
		if d.HasExpired(v) {
			d.Del(key)
			return nil
		}
	}

	return v
}

func (d *Dict) evictFirst() {
	for k := range d.items {
		delete(d.items, k)
		return
	}
}

func (d *Dict) evict() {
	switch config.EvictStrategy {
	case config.EvictFirst:
		d.evictFirst()
	// TODO: implement LRU and LFU eviction strategies
	default:
		d.evictFirst()
	}
}
