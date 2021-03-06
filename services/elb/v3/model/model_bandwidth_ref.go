package model

import (
	"encoding/json"

	"strings"
)

// 带宽信息id引用对象
type BandwidthRef struct {
	// 共享带宽的id

	Id string `json:"id"`
}

func (o BandwidthRef) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BandwidthRef struct{}"
	}

	return strings.Join([]string{"BandwidthRef", string(data)}, " ")
}
