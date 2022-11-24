package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTDA数据源详细配置内容
type IotdaContentRsp struct {

	// IoTDA实例Id
	IotdaInstanceId string `json:"iotda_instance_id"`

	// 在IoTDA实例中要配置的转发规则，包含推送数据的资源空间和产品列表
	Rules []IotdaContentRuleRsp `json:"rules"`
}

func (o IotdaContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IotdaContentRsp struct{}"
	}

	return strings.Join([]string{"IotdaContentRsp", string(data)}, " ")
}
