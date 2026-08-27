package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo 返回列表分页信息
type PageInfo struct {

	// **参数解释**： 偏移量。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,100000000]。 **默认取值**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每条页数。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,10000]。 **默认取值**： 不涉及
	PageSize *int32 `json:"page_size,omitempty"`

	// **参数解释**： 总条目数。 **约束限制**： 不涉及 **取值范围**： 取值范围[0,100000000]。 **默认取值**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
