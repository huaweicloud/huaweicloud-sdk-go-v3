package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSparkJobTemplatesResponse Response Object
type ListSparkJobTemplatesResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 返回的模板个数
	Count *int32 `json:"count,omitempty"`

	// 模板信息列表。
	Templates      *[]SparkJobTemplate `json:"templates,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListSparkJobTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListSparkJobTemplatesResponse", string(data)}, " ")
}
