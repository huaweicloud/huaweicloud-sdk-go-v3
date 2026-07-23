package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteDocRequestDto struct {

	// **参数解释：**  文档ID列表，系统生成的文档主键唯一标识列表，用于指定待批量删除的目标文档。 若不传入，则不执行任何删除操作。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Ids *[]string `json:"ids,omitempty"`

	// **参数解释：**  模型名称，用于指定待删除文档所属的数据模型名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ModelName *string `json:"model_name,omitempty"`

	// **参数解释：**  是否检查文档删除权限，用于控制删除前是否进行权限校验。  **约束限制：**  不涉及。  **取值范围：**  - true：检查删除权限，当前用户无权限的文档将被跳过删除。 - false：不检查删除权限，直接执行删除操作。  **默认取值：**  true。
	IsCheck *bool `json:"is_check,omitempty"`
}

func (o BatchDeleteDocRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDocRequestDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteDocRequestDto", string(data)}, " ")
}
