package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventSpecsRequest Request Object
type ListEventSpecsRequest struct {

	// 事件配置名称
	SpecName *string `json:"spec_name,omitempty"`

	// 事件类别
	Category *string `json:"category,omitempty"`

	// 事件级别
	Severity *string `json:"severity,omitempty"`

	// 事件源类别
	SourceType *string `json:"source_type,omitempty"`

	// 事件标签
	Tag *string `json:"tag,omitempty"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 1000
	Limit *string `json:"limit,omitempty"`
}

func (o ListEventSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSpecsRequest struct{}"
	}

	return strings.Join([]string{"ListEventSpecsRequest", string(data)}, " ")
}
