package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolume Volume信息。
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
