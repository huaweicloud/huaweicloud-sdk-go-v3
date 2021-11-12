package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCustomAuthorizerV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 自定义认证的ID

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
