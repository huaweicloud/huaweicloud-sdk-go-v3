package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 在IoTDA实例中要配置转发规则推送数据的资源空间和产品列表
type IotdaContentRuleReq struct {

	// IoTDA中的资源空间Id
	AppId string `json:"app_id"`

	// IoTDA中某资源空间Id下的产品列表
	Products []string `json:"products"`
}

func (o IotdaContentRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IotdaContentRuleReq struct{}"
	}

	return strings.Join([]string{"IotdaContentRuleReq", string(data)}, " ")
}
