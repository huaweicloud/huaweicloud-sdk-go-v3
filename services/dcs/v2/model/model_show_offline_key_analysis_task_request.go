package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOfflineKeyAnalysisTaskRequest Request Object
type ShowOfflineKeyAnalysisTaskRequest struct {

	// **参数解释**： 实例ID。可通过DCS控制台进入实例详情界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 任务ID。可通过DCS控制台进入离线全量key分析界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId string `json:"task_id"`
}

func (o ShowOfflineKeyAnalysisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOfflineKeyAnalysisTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowOfflineKeyAnalysisTaskRequest", string(data)}, " ")
}
