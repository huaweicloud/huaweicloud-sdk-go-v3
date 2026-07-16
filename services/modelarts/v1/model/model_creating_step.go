package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatingStep 批量创建信息。
type CreatingStep struct {

	// **参数解释**：超节点的步长。仅支持资源规格详情中包含的步长。 **取值范围**：不涉及。
	Step *int32 `json:"step,omitempty"`

	// **参数解释**：批量创建类型。 **取值范围**：可选值如下： - hyperinstance：超节点。
	Type *string `json:"type,omitempty"`
}

func (o CreatingStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatingStep struct{}"
	}

	return strings.Join([]string{"CreatingStep", string(data)}, " ")
}
