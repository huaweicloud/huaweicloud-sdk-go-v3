package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplateResponse Response Object
type UpdateTemplateResponse struct {

	// 模板主键ID
	Id *string `json:"id,omitempty"`

	// 模板名称
	TemplateName   *string `json:"template_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTemplateResponse", string(data)}, " ")
}
