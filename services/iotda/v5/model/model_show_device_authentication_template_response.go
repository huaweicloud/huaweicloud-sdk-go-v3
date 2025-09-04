package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeviceAuthenticationTemplateResponse Response Object
type ShowDeviceAuthenticationTemplateResponse struct {

	// 鉴权模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 鉴权模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 鉴权模板创建的时间。格式：yyyyMMdd'T'HHmmss'Z'，如：20151212T121212Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 鉴权模板最后一次修改的时间。格式：yyyyMMdd'T'HHmmss'Z'，如：20151212T121212Z。
	UpdateTime *string `json:"update_time,omitempty"`

	// 鉴权模板的描述信息
	Description *string `json:"description,omitempty"`

	// **参数说明**：鉴权模板状态 - ACTIVE：该鉴权模板为激活状态。 - INACTIVE：该鉴权模板为停用状态。
	Status *string `json:"status,omitempty"`

	TemplateBody   *AuthenticationTemplateBody `json:"template_body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowDeviceAuthenticationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeviceAuthenticationTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowDeviceAuthenticationTemplateResponse", string(data)}, " ")
}
