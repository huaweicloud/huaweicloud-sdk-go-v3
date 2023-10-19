package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConfigTemplateRequest Request Object
type DeleteConfigTemplateRequest struct {

	// 模板ID
	TemplateId string `json:"template_id"`
}

func (o DeleteConfigTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigTemplateRequest", string(data)}, " ")
}
