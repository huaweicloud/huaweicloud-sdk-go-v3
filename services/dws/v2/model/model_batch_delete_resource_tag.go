package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTag 标签详情。
type BatchDeleteResourceTag struct {

	// **参数解释**： 标签键。 **取值范围**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 标签值。 **取值范围**： 不涉及。
	Value string `json:"value"`
}

func (o BatchDeleteResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTag struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTag", string(data)}, " ")
}
