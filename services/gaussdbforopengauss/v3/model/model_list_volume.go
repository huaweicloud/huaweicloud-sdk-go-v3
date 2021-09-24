package model

import (
	"encoding/json"

	"strings"
)

// Volume信息。
type ListVolume struct {
	// 磁盘类型。

	Type string `json:"type"`
	// 磁盘大小。

	Size int32 `json:"size"`
}

func (o ListVolume) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListVolume struct{}"
	}

	return strings.Join([]string{"ListVolume", string(data)}, " ")
}
