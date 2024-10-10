package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListDDoSAttackEventRequestBody struct {

	// 开始时间（毫秒时间戳）
	StartTime string `json:"start_time"`

	// 结束时间（毫秒时间戳）
	EndTime string `json:"end_time"`

	// 偏移量
	Offset int32 `json:"offset"`

	// 限制条数，范围1-100
	Limit int32 `json:"limit"`

	// 高防ip
	Ip string `json:"ip"`

	// 攻击流量最小值
	AttackFlowLow *string `json:"attack_flow_low,omitempty"`

	// 攻击流量最大值
	AttackFlowUp *string `json:"attack_flow_up,omitempty"`

	// 攻击状态：attack-攻击; normal-结束攻击
	AttackStatus *string `json:"attack_status,omitempty"`
}

func (o ListDDoSAttackEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSAttackEventRequestBody struct{}"
	}

	return strings.Join([]string{"ListDDoSAttackEventRequestBody", string(data)}, " ")
}
