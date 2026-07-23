package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataCategoryDto struct {
	Categoryable *ObjectReferenceParamDto `json:"categoryable,omitempty"`

	Category *ObjectReferenceParamDto `json:"category,omitempty"`

	// **参数解释：**  更新者账号，用于记录执行本次分类添加操作的用户信息。 若不指定，默认使用当前调用者账号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  当前调用者账号。
	Modifier *string `json:"modifier,omitempty"`
}

func (o DataCategoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataCategoryDto struct{}"
	}

	return strings.Join([]string{"DataCategoryDto", string(data)}, " ")
}
