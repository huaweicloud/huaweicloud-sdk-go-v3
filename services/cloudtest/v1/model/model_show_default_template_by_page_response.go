package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultTemplateByPageResponse Response Object
type ShowDefaultTemplateByPageResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Data *BasePageInfoTemplate `json:"data,omitempty"`

	// 错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDefaultTemplateByPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultTemplateByPageResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultTemplateByPageResponse", string(data)}, " ")
}
