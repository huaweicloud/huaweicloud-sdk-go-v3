package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowComponentTemplateResponse Response Object
type ShowComponentTemplateResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 请求状态
	Success *bool `json:"success,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	Data           *TemplateInfo `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowComponentTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComponentTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowComponentTemplateResponse", string(data)}, " ")
}
