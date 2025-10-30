package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayoutLocation 设备位置信息
type LayoutLocation struct {

	// 机柜ID
	RackId *string `json:"rack_id,omitempty"`

	// 起始U位
	StartIndex *int32 `json:"start_index,omitempty"`

	// U位数
	Unit *int32 `json:"unit,omitempty"`
}

func (o LayoutLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutLocation struct{}"
	}

	return strings.Join([]string{"LayoutLocation", string(data)}, " ")
}
