package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadIdentitiesRequest Request Object
type ListWorkloadIdentitiesRequest struct {

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。
	Marker *string `json:"marker,omitempty"`
}

func (o ListWorkloadIdentitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadIdentitiesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadIdentitiesRequest", string(data)}, " ")
}
