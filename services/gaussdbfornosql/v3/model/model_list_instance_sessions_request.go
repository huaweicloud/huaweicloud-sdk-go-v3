package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceSessionsRequest Request Object
type ListInstanceSessionsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 节点ID。
	NodeId *string `json:"node_id,omitempty"`
}

func (o ListInstanceSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceSessionsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceSessionsRequest", string(data)}, " ")
}
