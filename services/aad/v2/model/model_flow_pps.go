package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlowPps struct {

	// 数据时间
	Utime *int64 `json:"utime,omitempty"`

	// 攻击包速率
	AttackPps *int32 `json:"attack_pps,omitempty"`

	// 正常包速率
	NormalPps *int32 `json:"normal_pps,omitempty"`
}

func (o FlowPps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowPps struct{}"
	}

	return strings.Join([]string{"FlowPps", string(data)}, " ")
}
