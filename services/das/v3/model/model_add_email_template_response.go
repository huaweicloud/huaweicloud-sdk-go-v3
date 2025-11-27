package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEmailTemplateResponse Response Object
type AddEmailTemplateResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 邮件模板ID
	TemplateId     *int32 `json:"template_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o AddEmailTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEmailTemplateResponse struct{}"
	}

	return strings.Join([]string{"AddEmailTemplateResponse", string(data)}, " ")
}
