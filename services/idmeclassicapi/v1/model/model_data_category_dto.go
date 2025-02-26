package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataCategoryDto struct {
	Categoryable *ObjectReferenceParamDto `json:"categoryable,omitempty"`

	Category *ObjectReferenceParamDto `json:"category,omitempty"`

	// **参数解释：**  修改人。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o DataCategoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataCategoryDto struct{}"
	}

	return strings.Join([]string{"DataCategoryDto", string(data)}, " ")
}
