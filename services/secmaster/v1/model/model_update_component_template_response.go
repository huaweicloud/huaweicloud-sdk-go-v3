package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComponentTemplateResponse Response Object
type UpdateComponentTemplateResponse struct {

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

func (o UpdateComponentTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateComponentTemplateResponse", string(data)}, " ")
}
