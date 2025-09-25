package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetResourceRequest Request Object
type ListAssetResourceRequest struct {

	// **参数解释**： 规格编码。 **约束限制**： 不涉及 **取值范围**： 长度为[1-64]个字符。 **默认取值**： 不涉及
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**： 资产的订购状态。 **约束限制**： 数组数量范围[0-10]，数组元素仅支持Normal|Freeze|Deleted。 **取值范围**： 不涉及 **默认取值**： 不涉及
	StatusList *[]string `json:"status_list,omitempty"`

	// **参数解释**： 限制量，单次查询总量。 **约束限制**： 不涉及 **取值范围**： 1-1000 **默认取值**： 100
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，查询起始偏移。 **约束限制**： 不涉及 **取值范围**： 0-100000000 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAssetResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetResourceRequest struct{}"
	}

	return strings.Join([]string{"ListAssetResourceRequest", string(data)}, " ")
}
