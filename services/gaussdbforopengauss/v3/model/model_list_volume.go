package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolume struct{}"
	}

	return strings.Join([]string{"ListVolume", string(data)}, " ")
}
