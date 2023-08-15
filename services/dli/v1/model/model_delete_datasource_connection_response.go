package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDatasourceConnectionResponse Response Object
type DeleteDatasourceConnectionResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，返回“Deleted”。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDatasourceConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatasourceConnectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteDatasourceConnectionResponse", string(data)}, " ")
}
