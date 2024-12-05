package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearInstanceSessionsRequest Request Object
type ClearInstanceSessionsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty"`
}

func (o ClearInstanceSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearInstanceSessionsRequest struct{}"
	}

	return strings.Join([]string{"ClearInstanceSessionsRequest", string(data)}, " ")
}
