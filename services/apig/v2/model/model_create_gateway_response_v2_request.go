package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateGatewayResponseV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`

	Body *ResponsesCreate `json:"body,omitempty"`
}

func (o CreateGatewayResponseV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGatewayResponseV2Request struct{}"
	}

	return strings.Join([]string{"CreateGatewayResponseV2Request", string(data)}, " ")
}
