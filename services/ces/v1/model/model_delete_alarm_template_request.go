package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmTemplateRequest Request Object
type DeleteAlarmTemplateRequest struct {

	// 需要删除的自定义告警模板ID。
	TemplateId string `json:"template_id"`
}

func (o DeleteAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmTemplateRequest", string(data)}, " ")
}
