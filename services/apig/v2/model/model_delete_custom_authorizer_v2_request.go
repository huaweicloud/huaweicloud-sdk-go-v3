package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomAuthorizerV2Request Request Object
type DeleteCustomAuthorizerV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 自定义认证的编号
	AuthorizerId string `json:"authorizer_id"`
}

func (o DeleteCustomAuthorizerV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomAuthorizerV2Request struct{}"
	}

	return strings.Join([]string{"DeleteCustomAuthorizerV2Request", string(data)}, " ")
}
