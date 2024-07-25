package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionUndoCheckOutDto struct {

	// **参数解释：**  主对象ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  更新者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionUndoCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionUndoCheckOutDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionUndoCheckOutDto", string(data)}, " ")
}
