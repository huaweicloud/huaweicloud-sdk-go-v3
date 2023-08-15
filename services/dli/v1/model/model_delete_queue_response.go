package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueResponse Response Object
type DeleteQueueResponse struct {

	// 请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueResponse struct{}"
	}

	return strings.Join([]string{"DeleteQueueResponse", string(data)}, " ")
}
