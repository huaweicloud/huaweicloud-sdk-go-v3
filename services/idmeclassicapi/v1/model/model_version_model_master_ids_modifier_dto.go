package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelMasterIdsModifierDto struct {

	// **参数解释：**  主对象ID列表，用于批量定位待删除的版本实例所属的主对象。  **约束限制：**  - 单次请求最多支持100个主对象ID。 - 每个主对象ID必须有效且存在对应M-V模型实例。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterIds []string `json:"masterIds"`

	// **参数解释：**  更新者账号，记录执行批量删除操作的用户标识。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelMasterIdsModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelMasterIdsModifierDto struct{}"
	}

	return strings.Join([]string{"VersionModelMasterIdsModifierDto", string(data)}, " ")
}
