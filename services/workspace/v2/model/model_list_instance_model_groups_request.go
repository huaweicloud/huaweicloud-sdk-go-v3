package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceModelGroupsRequest Request Object
type ListInstanceModelGroupsRequest struct {

	// Agent 实例 ID
	InstanceId string `json:"instance_id"`

	// 偏移量，从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceModelGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceModelGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceModelGroupsRequest", string(data)}, " ")
}
