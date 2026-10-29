package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessInfo **参数解释：** 执行记录信息。
type ProcessInfo struct {

	// **参数解释：** 执行时间，单位为秒。 **取值范围：** 不涉及。
	ExeTime *string `json:"exeTime,omitempty"`

	// **参数解释：** 执行时间，格式为%dm%ds，例如\"6m6s\"。 **取值范围：** 不涉及。
	ExecuteTime *string `json:"executeTime,omitempty"`

	// **参数解释：** 执行过程。 **取值范围：** 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释：** 任务执行的进度百分比。 **取值范围：** 不涉及。
	Process *string `json:"process,omitempty"`

	// **参数解释:** 任务执行状态。 **取值范围:** - quering：初始状态。 - running：执行中。 - success：检查成功。 - failed：检查失败。 - aborted：任务中止。
	Status *string `json:"status,omitempty"`
}

func (o ProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInfo struct{}"
	}

	return strings.Join([]string{"ProcessInfo", string(data)}, " ")
}
