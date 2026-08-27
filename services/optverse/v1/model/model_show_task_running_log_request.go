package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskRunningLogRequest Request Object
type ShowTaskRunningLogRequest struct {

	// **参数解释**： 演化任务标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	EvolveTaskId string `json:"evolve_task_id"`

	// **参数解释**： 算法的启动时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	StartByte *int64 `json:"start_byte,omitempty"`

	// **参数解释**： 算法的最后更新时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	EndByte *int64 `json:"end_byte,omitempty"`
}

func (o ShowTaskRunningLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskRunningLogRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskRunningLogRequest", string(data)}, " ")
}
