package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionUndoCheckOutDto struct {

	// **参数解释：**  M-V模型主对象ID，用于定位需要撤销检出的版本实例。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  撤销检出操作者账号，用于记录本次管理操作的责任人。由系统自动记录为当前调用用户，管理员操作时无需传入。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionUndoCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionUndoCheckOutDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionUndoCheckOutDto", string(data)}, " ")
}
