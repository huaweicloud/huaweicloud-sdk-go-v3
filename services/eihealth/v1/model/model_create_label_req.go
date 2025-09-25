package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLabelReq struct {

	// 标签名称
	Name string `json:"name"`

	// 标签描述
	Description *string `json:"description,omitempty"`

	// **参数解释**:  标签级别，用于前端颜色展示。   **约束限制**:  不涉及   **取值范围**:  整数，取值范围[1-6]   **默认取值**:  1
	Level *int32 `json:"level,omitempty"`
}

func (o CreateLabelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLabelReq struct{}"
	}

	return strings.Join([]string{"CreateLabelReq", string(data)}, " ")
}
