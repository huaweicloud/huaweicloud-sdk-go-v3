package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonthlyOperationReportsRequest Request Object
type ListMonthlyOperationReportsRequest struct {

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListMonthlyOperationReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonthlyOperationReportsRequest struct{}"
	}

	return strings.Join([]string{"ListMonthlyOperationReportsRequest", string(data)}, " ")
}
