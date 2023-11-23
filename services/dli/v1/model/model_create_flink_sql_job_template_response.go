package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSqlJobTemplateResponse Response Object
type CreateFlinkSqlJobTemplateResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容
	Message *string `json:"message,omitempty"`

	Template       *FlinkTemplate `json:"template,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateFlinkSqlJobTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobTemplateResponse", string(data)}, " ")
}
