package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteEmailTemplateRequestBody struct {

	// 邮件模板ID
	TemplateId int32 `json:"template_id"`
}

func (o DeleteEmailTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEmailTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteEmailTemplateRequestBody", string(data)}, " ")
}
