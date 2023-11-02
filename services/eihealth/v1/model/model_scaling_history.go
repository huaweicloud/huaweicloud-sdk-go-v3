package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScalingHistory struct {

	// 策略规则
	Rule string `json:"rule"`

	// 执行动作
	Action string `json:"action"`

	// 伸缩前节点数
	CountBeforeScale int32 `json:"count_before_scale"`

	// 伸缩后节点数
	CountAfterScale int32 `json:"count_after_scale"`

	// 执行状态
	State string `json:"state"`

	// 执行时间
	Time string `json:"time"`
}

func (o ScalingHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingHistory struct{}"
	}

	return strings.Join([]string{"ScalingHistory", string(data)}, " ")
}
