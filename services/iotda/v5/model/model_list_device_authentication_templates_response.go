package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeviceAuthenticationTemplatesResponse Response Object
type ListDeviceAuthenticationTemplatesResponse struct {

	// **参数说明**：鉴权模板详情列表。
	Templates *[]AuthenticationTemplateSimple `json:"templates,omitempty"`

	Page           *Page `json:"page,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListDeviceAuthenticationTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeviceAuthenticationTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListDeviceAuthenticationTemplatesResponse", string(data)}, " ")
}
