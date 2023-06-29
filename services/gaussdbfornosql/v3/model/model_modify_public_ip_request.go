package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyPublicIpRequest Request Object
type ModifyPublicIpRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 实例节点ID。
	NodeId string `json:"node_id"`

	Body *ModifyPublicIpRequestBody `json:"body,omitempty"`
}

func (o ModifyPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPublicIpRequest struct{}"
	}

	return strings.Join([]string{"ModifyPublicIpRequest", string(data)}, " ")
}
