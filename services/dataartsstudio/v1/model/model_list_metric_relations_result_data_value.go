package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricRelationsResultDataValue value，统一的返回结果的外层数据结构。
type ListMetricRelationsResultDataValue struct {

	// 所有的业务指标信息。
	All *[]interface{} `json:"all,omitempty"`

	// 指标关联。
	Links *interface{} `json:"links,omitempty"`

	// 分组。
	Groups *interface{} `json:"groups,omitempty"`

	// 总数。
	Total *int32 `json:"total,omitempty"`
}

func (o ListMetricRelationsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricRelationsResultDataValue struct{}"
	}

	return strings.Join([]string{"ListMetricRelationsResultDataValue", string(data)}, " ")
}
