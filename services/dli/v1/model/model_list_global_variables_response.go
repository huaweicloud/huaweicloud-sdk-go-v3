package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalVariablesResponse Response Object
type ListGlobalVariablesResponse struct {

	// 请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 全局变量
	GlobalVars *[]GlobalVariable `json:"global_vars,omitempty"`

	// 全局变量总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGlobalVariablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalVariablesResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalVariablesResponse", string(data)}, " ")
}
