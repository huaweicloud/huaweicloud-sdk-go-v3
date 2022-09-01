package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThrottleApiBindingCreate struct {

	// 流控策略编号
	StrategyId string `json:"strategy_id" xml:"strategy_id"`

	// API的发布记录编号
	PublishIds []string `json:"publish_ids" xml:"publish_ids"`
}

func (o ThrottleApiBindingCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThrottleApiBindingCreate struct{}"
	}

	return strings.Join([]string{"ThrottleApiBindingCreate", string(data)}, " ")
}
