package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 机柜信息
type RackInfo struct {

	// 机柜功率，单位：w
	Power *int32 `json:"power,omitempty"`

	// 机柜尺寸，如100\\*200\\*200，单位：cm
	Size *string `json:"size,omitempty"`

	// 是否有机柜锁。
	HasLock *bool `json:"has_lock,omitempty"`
}

func (o RackInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RackInfo struct{}"
	}

	return strings.Join([]string{"RackInfo", string(data)}, " ")
}
