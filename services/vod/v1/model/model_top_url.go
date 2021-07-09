package model

import (
	"encoding/json"

	"strings"
)

type TopUrl struct {
	// 总播放次数。

	Value *int64 `json:"value,omitempty"`
	// 媒资ID。

	AssetId *string `json:"asset_id,omitempty"`
	// 媒资名称。

	Title *string `json:"title,omitempty"`
	// 媒资时长。  单位：秒。

	Duration *int32 `json:"duration,omitempty"`
	// 媒资原始大小。  单位：字节。

	Size *int64 `json:"size,omitempty"`
}

func (o TopUrl) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TopUrl struct{}"
	}

	return strings.Join([]string{"TopUrl", string(data)}, " ")
}
