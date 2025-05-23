package logic

import (
	"encoding/json"
	"sync/atomic"
)

type Stats struct {
	// 分发总消息数
	DispatchTotal int64 `json:"DispatchTotal"`
	// 分发丢弃消息数
	DispatchFail int64 `json:"DispatchFail"`
	// 推送失败次数
	PushFail int64 `json:"PushFail"`
}

var (
	G_stats *Stats
)

func InitStats() (err error) {
	G_stats = &Stats{}
	return
}

// 增加分发消息数
func DispatchTotal_INCR(batchSize int64) {
	atomic.AddInt64(&G_stats.DispatchTotal, batchSize)
}

// 增加分发丢弃消息数
func DispatchFail_INCR(batchSize int64) {
	atomic.AddInt64(&G_stats.DispatchFail, batchSize)
}

// 增加推送失败次数
func PushFail_INCR() {
	atomic.AddInt64(&G_stats.PushFail, 1)
}

// 获取统计信息
func (stats *Stats) Dump() (data []byte, err error) {
	return json.Marshal(G_stats)
}
