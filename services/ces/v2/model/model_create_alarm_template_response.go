package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlarmTemplateResponse Response Object
type CreateAlarmTemplateResponse struct {

	// 告警模板的ID，以at开头，后跟字母、数字，长度最长为64
	TemplateId     *string `json:"template_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateResponse", string(data)}, " ")
}
