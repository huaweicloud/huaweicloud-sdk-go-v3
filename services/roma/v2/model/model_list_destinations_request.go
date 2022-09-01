package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDestinationsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 规则ID
	RuleId string `json:"rule_id" xml:"rule_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListDestinationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDestinationsRequest struct{}"
	}

	return strings.Join([]string{"ListDestinationsRequest", string(data)}, " ")
}
