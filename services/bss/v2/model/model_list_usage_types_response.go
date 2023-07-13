package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsageTypesResponse Response Object
type ListUsageTypesResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 使用量类型列表，具体请参见表3。
	UsageTypes     *[]UsageType `json:"usage_types,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListUsageTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsageTypesResponse struct{}"
	}

	return strings.Join([]string{"ListUsageTypesResponse", string(data)}, " ")
}
