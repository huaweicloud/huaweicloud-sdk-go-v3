package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDdmInstanceNameRequestBody struct {

	// **参数解释**：  实例名称。  **约束限制**：  - 长度为4-64个字符。 - 必须以字母开头。 - 可以包含字母、数字、中划线、下划线，不能包含其它特殊字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Name string `json:"name"`
}

func (o UpdateDdmInstanceNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDdmInstanceNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDdmInstanceNameRequestBody", string(data)}, " ")
}
