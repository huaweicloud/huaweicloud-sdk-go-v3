package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecCode **参数解释**：资源规格。 **取值范围**：不涉及。
type SpecCode struct {
}

func (o SpecCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecCode struct{}"
	}

	return strings.Join([]string{"SpecCode", string(data)}, " ")
}
