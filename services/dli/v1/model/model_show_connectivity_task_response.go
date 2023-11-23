package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectivityTaskResponse Response Object
type ShowConnectivityTaskResponse struct {

	// 请求发送是否成功。“true”表示请求发送成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 连通性测试结果
	Connectivity   *string `json:"connectivity,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConnectivityTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectivityTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowConnectivityTaskResponse", string(data)}, " ")
}
