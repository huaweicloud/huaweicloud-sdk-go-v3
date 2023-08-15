package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomAuthorizerV2Request Request Object
type CreateCustomAuthorizerV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *AuthorizerCreate `json:"body,omitempty"`
}

func (o CreateCustomAuthorizerV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomAuthorizerV2Request struct{}"
	}

	return strings.Join([]string{"CreateCustomAuthorizerV2Request", string(data)}, " ")
}
