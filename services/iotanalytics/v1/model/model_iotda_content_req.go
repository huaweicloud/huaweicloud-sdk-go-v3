package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTDA数据源详细配置内容
type IotdaContentReq struct {

	// IoTDA实例Id
	IotdaInstanceId string `json:"iotda_instance_id"`

	// 在IoTDA实例中要配置转发规则推送数据的资源空间和产品列表
	Rules []IotdaContentRuleReq `json:"rules"`
}

func (o IotdaContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IotdaContentReq struct{}"
	}

	return strings.Join([]string{"IotdaContentReq", string(data)}, " ")
}
