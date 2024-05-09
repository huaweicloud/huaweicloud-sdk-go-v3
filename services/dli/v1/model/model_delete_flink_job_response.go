package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlinkJobResponse Response Object
type DeleteFlinkJobResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *string `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlinkJobResponse", string(data)}, " ")
}
