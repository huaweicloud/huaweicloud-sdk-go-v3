package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConnectivityTaskResponse Response Object
type CreateConnectivityTaskResponse struct {

	// 请求发送是否成功。“true”表示请求发送成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 请求id
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateConnectivityTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectivityTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateConnectivityTaskResponse", string(data)}, " ")
}
