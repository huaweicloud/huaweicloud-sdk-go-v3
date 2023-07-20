package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkTemplateResponse Response Object
type CreateFlinkTemplateResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容
	Message *string `json:"message,omitempty"`

	Template       *FlinkTemplate `json:"template,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateFlinkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkTemplateResponse", string(data)}, " ")
}
