package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComponentTemplateResponse Response Object
type DeleteComponentTemplateResponse struct {

	// **参数解释**: 响应的返回码 **约束限制**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 响应的错误信息 **约束限制**: 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**: 是否响应成功 **约束限制**: 不涉及
	Success *bool `json:"success,omitempty"`

	// **参数解释**: 请求id **约束限制**: 不涉及
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteComponentTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComponentTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteComponentTemplateResponse", string(data)}, " ")
}
