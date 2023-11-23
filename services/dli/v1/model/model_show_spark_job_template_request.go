package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobTemplateRequest Request Object
type ShowSparkJobTemplateRequest struct {
	TemplateId string `json:"template_id"`
}

func (o ShowSparkJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobTemplateRequest", string(data)}, " ")
}
