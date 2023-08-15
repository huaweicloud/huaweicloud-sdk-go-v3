package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmTemplateRequest Request Object
type UpdateAlarmTemplateRequest struct {

	// 需要更新的自定义告警模板ID。
	TemplateId string `json:"template_id"`

	Body *UpdateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmTemplateRequest", string(data)}, " ")
}
