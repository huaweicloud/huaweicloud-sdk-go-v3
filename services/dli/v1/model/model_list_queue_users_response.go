package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueueUsersResponse Response Object
type ListQueueUsersResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 队列名称。
	QueueName *string `json:"queue_name,omitempty"`

	// 有权限使用该队列的用户及其对应的权限数组。
	Privileges     *[]QueueUserPrivilege `json:"privileges,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListQueueUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueueUsersResponse struct{}"
	}

	return strings.Join([]string{"ListQueueUsersResponse", string(data)}, " ")
}
