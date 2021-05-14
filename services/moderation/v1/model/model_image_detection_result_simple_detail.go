package model

import (
	"encoding/json"

	"strings"
)

// 反黄，暴恐，广告检测详情
type ImageDetectionResultSimpleDetail struct {
	Confidence *float32 `json:"confidence,omitempty"`

	Label *string `json:"label,omitempty"`
}

func (o ImageDetectionResultSimpleDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageDetectionResultSimpleDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultSimpleDetail", string(data)}, " ")
}
