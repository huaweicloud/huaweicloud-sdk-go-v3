package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMyActionTemplateResponse Response Object
type UpdateMyActionTemplateResponse struct {

	// 算子模板名称。
	TemplateName *string `json:"template_name,omitempty"`

	// 算子模板创建的时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 版本号。
	Version *string `json:"version,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMyActionTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMyActionTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateMyActionTemplateResponse", string(data)}, " ")
}
