package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeContentReq IoTDA数据源详细配置内容
type EdgeContentReq struct {

	// edge实例Id
	IotdaInstanceId *string `json:"iotda_instance_id,omitempty"`

	// 在edge实例中要配置转发规则推送数据的资源空间和产品列表
	Rules []EdgeContentRuleReq `json:"rules"`
}

func (o EdgeContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeContentReq struct{}"
	}

	return strings.Join([]string{"EdgeContentReq", string(data)}, " ")
}
