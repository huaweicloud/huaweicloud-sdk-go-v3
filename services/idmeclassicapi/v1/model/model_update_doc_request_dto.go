package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDocRequestDto struct {

	// **参数解释：**  KooPage文档ID，用于唯一标识待更新的结构化文档。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	DocumentId string `json:"document_id"`

	// **参数解释：**  文档标题，用于指定更新后的文档名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  实例ID，用于通过关联的数据模型实例定位待更新的文档。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释：**  更新者账号，用于记录执行本次更新操作的用户信息。 若不指定，默认使用当前调用者账号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  当前调用者账号。
	Modifier *string `json:"modifier,omitempty"`
}

func (o UpdateDocRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDocRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateDocRequestDto", string(data)}, " ")
}
