package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultCount struct {

	// **参数解释**：  项目下的版本数量。 **取值范围**：  不涉及。
	Count *int32 `json:"count,omitempty"`
}

func (o ResultCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultCount struct{}"
	}

	return strings.Join([]string{"ResultCount", string(data)}, " ")
}
