package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Failure struct {
	// API请求路径

	Path *string `json:"path,omitempty"`
	// 导入失败的错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
	// API请求方法

	Method *string `json:"method,omitempty"`
	// 导入失败的错误码

	ErrorCode *string `json:"error_code,omitempty"`
}

func (o Failure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Failure struct{}"
	}

	return strings.Join([]string{"Failure", string(data)}, " ")
}
