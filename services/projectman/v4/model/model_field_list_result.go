package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FieldListResult struct {

	// **参数解释**： 字段列表。
	Data *[]FieldEntity `json:"data,omitempty"`

	// **参数解释**： 项目内字段总数。 **取值范围**： 不涉及
	Total *int32 `json:"total,omitempty"`
}

func (o FieldListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldListResult struct{}"
	}

	return strings.Join([]string{"FieldListResult", string(data)}, " ")
}
