package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityAssignedQueuesRequest Request Object
type ListSecurityAssignedQueuesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 队列类型，MRS、DLI。
	Type *string `json:"type,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSecurityAssignedQueuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityAssignedQueuesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityAssignedQueuesRequest", string(data)}, " ")
}
