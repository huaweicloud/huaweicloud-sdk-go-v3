package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResendEmailReq struct {

	// 邮件模板ID。
	TemplateId *string `json:"template_id,omitempty"`
}

func (o ResendEmailReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResendEmailReq struct{}"
	}

	return strings.Join([]string{"ResendEmailReq", string(data)}, " ")
}
