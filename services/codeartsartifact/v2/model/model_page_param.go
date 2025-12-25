package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageParam struct {

	// **参数解释**: 页码。 **约束限制**: 不涉及。 **取值范围**: 最小值1。 **默认取值**: 1
	PageNo *int32 `json:"page_no,omitempty"`

	// **参数解释**: 每页大小。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值100。 **默认取值**: 10
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o PageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageParam struct{}"
	}

	return strings.Join([]string{"PageParam", string(data)}, " ")
}
