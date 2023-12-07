package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobAuthInfoResponse Response Object
type DeleteJobAuthInfoResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteJobAuthInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobAuthInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobAuthInfoResponse", string(data)}, " ")
}
