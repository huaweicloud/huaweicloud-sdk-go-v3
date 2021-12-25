package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDestinationRequestBody struct {
	// 操作类型，枚举值:0-目标端为本ROMA实例内MQS，； 7-目标端为设备

	DestinationType int32 `json:"destination_type"`
	// 应用ID，目标端为0时需明确对方的APP_ID

	AppId *string `json:"app_id,omitempty"`
	// 目标数据源名称

	DestinationName *string `json:"destination_name,omitempty"`
	// 目标数据源主题，从MQS服务中获取已有topic

	Topic string `json:"topic"`
	// 目标端数据源服务，连接地址

	Server *string `json:"server,omitempty"`
	// 目标端数据源token

	Token *string `json:"token,omitempty"`
	// 目标数据源标签

	Tag *string `json:"tag,omitempty"`
	// 目标端数据源MQS的SASL字段是否需要支持SSL加密

	MqsSaslSsl *bool `json:"mqs_sasl_ssl,omitempty"`
	// 目标数据源用户名

	UserName *string `json:"user_name,omitempty"`
	// 目标数据源密码

	Password *string `json:"password,omitempty"`
}

func (o CreateDestinationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDestinationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDestinationRequestBody", string(data)}, " ")
}
