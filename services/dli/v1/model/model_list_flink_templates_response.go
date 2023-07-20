package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlinkTemplatesResponse Response Object
type ListFlinkTemplatesResponse struct {

	// 请求是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	TemplateList   *FlinkTemplateList `json:"template_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListFlinkTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListFlinkTemplatesResponse", string(data)}, " ")
}
