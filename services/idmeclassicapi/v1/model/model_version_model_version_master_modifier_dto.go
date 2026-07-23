package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionMasterModifierDto struct {

	// **参数解释：**  主对象ID，用于定位待删除的版本实例所属的主对象。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  更新者账号，记录执行删除操作的用户标识。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  版本对象，用于指定待删除的版本标识。若未指定，则默认删除最新分支的最新版本。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelVersionMasterModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionMasterModifierDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionMasterModifierDto", string(data)}, " ")
}
