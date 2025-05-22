package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadStatus **参数解释**： 资源管理状态请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadStatus struct {

	// **参数解释**： 开关。 **约束限制**： 不涉及。 **取值范围**： on：开启 off：关闭 **默认取值**： 不涉及。
	WorkloadSwitch string `json:"workload_switch"`

	// **参数解释**： 最大并发数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxConcurrencyNum *string `json:"max_concurrency_num,omitempty"`
}

func (o WorkloadStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadStatus struct{}"
	}

	return strings.Join([]string{"WorkloadStatus", string(data)}, " ")
}
