package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlowBps struct {

	// 数据时间
	Utime *int64 `json:"utime,omitempty"`

	// 攻击流量
	AttackBps *int32 `json:"attack_bps,omitempty"`

	// 正常流量
	NormalBps *int32 `json:"normal_bps,omitempty"`
}

func (o FlowBps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowBps struct{}"
	}

	return strings.Join([]string{"FlowBps", string(data)}, " ")
}
