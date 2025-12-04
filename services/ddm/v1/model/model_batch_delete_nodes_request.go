package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteNodesRequest Request Object
type BatchDeleteNodesRequest struct {

	// 实例 ID。
	InstanceId string `json:"instance_id"`

	Body *ReduceNodeOpenRequest `json:"body,omitempty"`
}

func (o BatchDeleteNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteNodesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteNodesRequest", string(data)}, " ")
}
