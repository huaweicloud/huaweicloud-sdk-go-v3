package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopEvolveTaskRequest Request Object
type StopEvolveTaskRequest struct {

	// **参数解释**： 演化任务标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	EvolveTaskId string `json:"evolve_task_id"`
}

func (o StopEvolveTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopEvolveTaskRequest struct{}"
	}

	return strings.Join([]string{"StopEvolveTaskRequest", string(data)}, " ")
}
