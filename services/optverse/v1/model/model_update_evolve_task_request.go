package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEvolveTaskRequest Request Object
type UpdateEvolveTaskRequest struct {

	// **参数解释**： 演化任务标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	EvolveTaskId string `json:"evolve_task_id"`

	Body *EvolveTaskCreateReq `json:"body,omitempty"`
}

func (o UpdateEvolveTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEvolveTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateEvolveTaskRequest", string(data)}, " ")
}
