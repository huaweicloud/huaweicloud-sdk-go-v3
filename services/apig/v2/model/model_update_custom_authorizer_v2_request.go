package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomAuthorizerV2Request Request Object
type UpdateCustomAuthorizerV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 自定义认证的编号
	AuthorizerId string `json:"authorizer_id"`

	Body *AuthorizerCreate `json:"body,omitempty"`
}

func (o UpdateCustomAuthorizerV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomAuthorizerV2Request struct{}"
	}

	return strings.Join([]string{"UpdateCustomAuthorizerV2Request", string(data)}, " ")
}
