package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateBinlogRequestV3Body 开启/关闭Binlog参数体。
type OperateBinlogRequestV3Body struct {

	// **参数解释**：  Binlog日志开关状态。  **约束限制**：  不涉及。  **取值范围**：  - ON：表示开启。 - OFF：表示关闭。  **默认取值**：  不涉及。
	SwitchStatus string `json:"switch_status"`

	// **参数解释**：  Binlog日志保存的天数。  **约束限制**：  - switch_status为ON时必选。 - switch_status为OFF时不可选。  **取值范围**：  1~180。  **默认取值**：  不涉及。
	KeepDays *int32 `json:"keep_days,omitempty"`
}

func (o OperateBinlogRequestV3Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateBinlogRequestV3Body struct{}"
	}

	return strings.Join([]string{"OperateBinlogRequestV3Body", string(data)}, " ")
}
