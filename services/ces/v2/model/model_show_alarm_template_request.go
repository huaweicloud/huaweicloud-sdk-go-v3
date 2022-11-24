package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAlarmTemplateRequest struct {

	// 告警模板的ID，以at开头，后跟字母、数字，长度最长为64
	TemplateId string `json:"template_id"`
}

func (o ShowAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmTemplateRequest", string(data)}, " ")
}
