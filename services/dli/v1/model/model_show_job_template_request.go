package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobTemplateRequest Request Object
type ShowJobTemplateRequest struct {
	TemplateId string `json:"template_id"`
}

func (o ShowJobTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowJobTemplateRequest", string(data)}, " ")
}
