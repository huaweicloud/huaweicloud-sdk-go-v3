package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePartitionCountRequest Request Object
type UpdatePartitionCountRequest struct {

	// 需要变更分区数量的通道名称。
	StreamName string `json:"stream_name"`

	Body *UpdatePartitionCountReq `json:"body,omitempty"`
}

func (o UpdatePartitionCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePartitionCountRequest struct{}"
	}

	return strings.Join([]string{"UpdatePartitionCountRequest", string(data)}, " ")
}
