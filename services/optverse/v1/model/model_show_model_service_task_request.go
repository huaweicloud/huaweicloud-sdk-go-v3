package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelServiceTaskRequest Request Object
type ShowModelServiceTaskRequest struct {

	// **参数解释**： 模型服务ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	ServiceId string `json:"service_id"`

	// **参数解释**： 任务ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	TaskId string `json:"task_id"`
}

func (o ShowModelServiceTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelServiceTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowModelServiceTaskRequest", string(data)}, " ")
}
