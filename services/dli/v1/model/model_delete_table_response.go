package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableResponse Response Object
type DeleteTableResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 删除表作业的时候类型，是同步删除还是异步删除
	JobMode        *string `json:"job_mode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableResponse", string(data)}, " ")
}
