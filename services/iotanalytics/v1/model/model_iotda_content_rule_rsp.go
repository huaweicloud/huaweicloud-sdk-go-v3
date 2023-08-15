package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IotdaContentRuleRsp 在IoTDA实例中要配置转发规则推送数据的资源空间和产品列表
type IotdaContentRuleRsp struct {

	// IoTDA中的资源空间Id
	AppId string `json:"app_id"`

	// IoTDA中某资源空间Id下的产品列表
	Products []string `json:"products"`

	// IoTDA中rule_id和action_id列表
	RuleActions []IotdaRuleAction `json:"rule_actions"`
}

func (o IotdaContentRuleRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IotdaContentRuleRsp struct{}"
	}

	return strings.Join([]string{"IotdaContentRuleRsp", string(data)}, " ")
}
