package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RpoAndRTO信息体
type RpoAndRtoInfo struct {
	// 检查点

	CheckPoint *string `json:"check_point,omitempty"`
	// 延迟

	Delay *string `json:"delay,omitempty"`
	// gtid集合

	GtidSet *string `json:"gtid_set,omitempty"`
	// 当前时间 ，格式为“yyyy-MM-dd HH:mm:ss”

	Time *string `json:"time,omitempty"`
}

func (o RpoAndRtoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RpoAndRtoInfo struct{}"
	}

	return strings.Join([]string{"RpoAndRtoInfo", string(data)}, " ")
}
