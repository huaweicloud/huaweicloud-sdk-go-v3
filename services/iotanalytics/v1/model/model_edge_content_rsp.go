package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTDA数据源详细配置内容
type EdgeContentRsp struct {
	// Edge实例Id

	IotdaInstanceId string `json:"iotda_instance_id"`
	// 在Edge实例中要配置的转发规则，包含推送数据的资源空间和产品列表

	Rules []EdgeContentRuleRsp `json:"rules"`
}

func (o EdgeContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeContentRsp struct{}"
	}

	return strings.Join([]string{"EdgeContentRsp", string(data)}, " ")
}
