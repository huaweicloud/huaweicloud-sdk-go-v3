package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValuesItem **参数解释：** 版本对象。 **取值范围：** 不涉及。
type ValuesItem struct {

	// **参数解释：** 所有版本列表。 **取值范围：** 不涉及。
	Values *[]ListApiVersionsItem `json:"values,omitempty"`
}

func (o ValuesItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValuesItem struct{}"
	}

	return strings.Join([]string{"ValuesItem", string(data)}, " ")
}
