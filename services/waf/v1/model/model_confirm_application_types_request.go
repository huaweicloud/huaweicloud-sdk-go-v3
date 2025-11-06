package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmApplicationTypesRequest Request Object
type ConfirmApplicationTypesRequest struct {

	// **参数解释：** 分页查询的起始位置，表示从第几条记录开始返回（从1开始计数）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 1
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 分页查询的单页返回数量，控制每次请求返回的记录条数。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ConfirmApplicationTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmApplicationTypesRequest struct{}"
	}

	return strings.Join([]string{"ConfirmApplicationTypesRequest", string(data)}, " ")
}
