package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCustomAuthorizerV2Request struct {
	// 实例ID

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
