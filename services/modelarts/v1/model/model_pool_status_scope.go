package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolStatusScope struct {

	// **参数解释**：资源池的业务类型。 **取值范围**：可选值如下： - Train：训练任务。 - Infer：推理任务。 - Notebook：Notebook作业。
	ScopeType string `json:"scopeType"`

	// **参数解释**：资源池业务类型状态。 **取值范围**：可选值如下： - Enabling：启动中。 - Enabled：已启动。 - Disabling：关闭中。 - Disabled：已关闭。
	State string `json:"state"`
}

func (o PoolStatusScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatusScope struct{}"
	}

	return strings.Join([]string{"PoolStatusScope", string(data)}, " ")
}
