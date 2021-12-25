package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMqsInstanceTopicsResponse struct {
	// topic的总数。

	Total *int32 `json:"total,omitempty"`
	// 当前查询的topic数量。

	Size *int32 `json:"size,omitempty"`
	// 允许操作的权限。

	Permissions *[]string `json:"permissions,omitempty"`
	// Topic列表。

	Topics *[]ListInstanceTopicsRespTopics `json:"topics,omitempty"`
	// 权限列表。

	Policies *[]ListInstanceTopicsRespPolicies `json:"policies,omitempty"`
	// 剩余分区数。

	RemainPartitions *int32 `json:"remain_partitions,omitempty"`
	// 分区总数。

	MaxPartitions  *int32 `json:"max_partitions,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMqsInstanceTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMqsInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListMqsInstanceTopicsResponse", string(data)}, " ")
}
