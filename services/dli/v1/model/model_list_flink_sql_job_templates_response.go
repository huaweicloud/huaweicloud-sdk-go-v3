package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkSqlJobTemplatesResponse Response Object
type ListFlinkSqlJobTemplatesResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	TemplateList   *FlinkSqlJobTemplateList `json:"template_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListFlinkSqlJobTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkSqlJobTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkSqlJobTemplatesResponse", string(data)}, " ")
}
