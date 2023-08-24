package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAimMsgTemplateResponse Response Object
type CreateAimMsgTemplateResponse struct {

	// 模板ID。
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称。
	TemplateName   *string `json:"template_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAimMsgTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAimMsgTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateAimMsgTemplateResponse", string(data)}, " ")
}
