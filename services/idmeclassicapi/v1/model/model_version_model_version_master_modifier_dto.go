package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionMasterModifierDto struct {

	// **参数解释：**  父模型ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  修改人。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  版本对象。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version *string `json:"version,omitempty"`
}

func (o VersionModelVersionMasterModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionMasterModifierDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionMasterModifierDto", string(data)}, " ")
}
