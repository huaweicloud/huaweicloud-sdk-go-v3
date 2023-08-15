package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckConnectionResponse Response Object
type CheckConnectionResponse struct {

	// 请求发送是否成功。“true”表示请求发送成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 请求id
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckConnectionResponse struct{}"
	}

	return strings.Join([]string{"CheckConnectionResponse", string(data)}, " ")
}
