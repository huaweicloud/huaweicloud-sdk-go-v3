package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecGuardListRequest Request Object
type ListSecGuardListRequest struct {

	// **参数解释**: 查询多少天内的任务。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 1。
	Date *int32 `json:"date,omitempty"`

	// **参数解释**: 分页查询的页数。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值2147483647。 **默认取值**: 1
	PageNo *int32 `json:"page_no,omitempty"`

	// **参数解释**: 分页查询的每页数据量。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值100。 **默认取值**: 10
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListSecGuardListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecGuardListRequest struct{}"
	}

	return strings.Join([]string{"ListSecGuardListRequest", string(data)}, " ")
}
