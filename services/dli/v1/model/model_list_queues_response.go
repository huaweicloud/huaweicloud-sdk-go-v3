package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuesResponse Response Object
type ListQueuesResponse struct {

	// 参数解释: 请求执行是否成功。“true”表示请求执行成功 示例: true 约束限制:  无 取值范围: 无 默认取值: 无
	IsSuccess *bool `json:"is_success,omitempty"`

	// 参数解释: 系统提示信息，执行成功时，信息可能为空 示例: test 约束限制:  无 取值范围: 无 默认取值: 无
	Message *string `json:"message,omitempty"`

	// 队列信息。
	Queues         *[]Queue `json:"queues,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesResponse struct{}"
	}

	return strings.Join([]string{"ListQueuesResponse", string(data)}, " ")
}
