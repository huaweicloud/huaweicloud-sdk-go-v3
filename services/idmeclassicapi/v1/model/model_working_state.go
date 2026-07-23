package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkingState struct {

	// **参数解释：**  工作状态别名，用于系统内部的状态标。  **取值范围：**  不涉及。
	Alias *string `json:"alias,omitempty"`

	// **参数解释：**  工作状态中文名称。  **取值范围：**  不涉及。
	CnName *string `json:"cnName,omitempty"`

	// **参数解释：**  工作状态编码。  **取值范围：**  不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释：**  工作状态英文名称。  **取值范围：**  不涉及。
	EnName *string `json:"enName,omitempty"`
}

func (o WorkingState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkingState struct{}"
	}

	return strings.Join([]string{"WorkingState", string(data)}, " ")
}
