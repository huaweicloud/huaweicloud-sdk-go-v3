package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 维护时间窗结构体
type MaintainWindowsEntity struct {
	// 序号。

	Seq *int32 `json:"seq,omitempty"`
	// 是否为默认时间段。

	Default *bool `json:"default,omitempty"`
	// 维护时间窗开始时间

	Begin *string `json:"begin,omitempty"`
	// 维护时间窗结束时间。

	End *string `json:"end,omitempty"`
}

func (o MaintainWindowsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaintainWindowsEntity struct{}"
	}

	return strings.Join([]string{"MaintainWindowsEntity", string(data)}, " ")
}
