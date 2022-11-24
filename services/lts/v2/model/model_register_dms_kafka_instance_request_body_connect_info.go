package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// kafka连接信息。购买kafka时，SASL_SSL是否配置，如果有，则需要填写账号密码
type RegisterDmsKafkaInstanceRequestBodyConnectInfo struct {

	// 账号
	UserName *string `json:"user_name,omitempty"`

	// 密码
	Pwd *string `json:"pwd,omitempty"`
}

func (o RegisterDmsKafkaInstanceRequestBodyConnectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDmsKafkaInstanceRequestBodyConnectInfo struct{}"
	}

	return strings.Join([]string{"RegisterDmsKafkaInstanceRequestBodyConnectInfo", string(data)}, " ")
}
