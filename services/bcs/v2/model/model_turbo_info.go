package model

import (
	"encoding/json"

	"strings"
)

// 极速文件存储卷信息
type TurboInfo struct {
	// 共享方式

	ShareType string `json:"share_type"`
	// 类型

	Type string `json:"type"`
	// 可用区

	AvailableZone string `json:"available_zone"`
	// 规格

	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o TurboInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TurboInfo struct{}"
	}

	return strings.Join([]string{"TurboInfo", string(data)}, " ")
}
