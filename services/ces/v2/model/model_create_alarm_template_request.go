package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlarmTemplateRequest Request Object
type CreateAlarmTemplateRequest struct {
	Body *CreateAlarmTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateRequest", string(data)}, " ")
}
