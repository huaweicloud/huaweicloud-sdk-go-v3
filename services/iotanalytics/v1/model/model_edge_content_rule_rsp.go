package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeContentRuleRsp 在Edge实例中要配置转发规则推送数据的资源空间和产品列表
type EdgeContentRuleRsp struct {

	// Edge中的资源空间Id
	AppId string `json:"app_id"`

	// Edge中某资源空间Id下的产品列表
	Products []string `json:"products"`

	// Edge中rule_id和action_id列表
	RuleActions []IotdaRuleAction `json:"rule_actions"`
}

func (o EdgeContentRuleRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeContentRuleRsp struct{}"
	}

	return strings.Join([]string{"EdgeContentRuleRsp", string(data)}, " ")
}
