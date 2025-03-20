package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlaCustomizedTemplateRequest Request Object
type ShowSlaCustomizedTemplateRequest struct {

	// 模板ID
	TemplateId string `json:"template_id"`
}

func (o ShowSlaCustomizedTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlaCustomizedTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowSlaCustomizedTemplateRequest", string(data)}, " ")
}
