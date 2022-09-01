package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Failure struct {

	// API请求路径
	Path *string `json:"path,omitempty" xml:"path"`

	// 导入失败的错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// API请求方法
	Method *string `json:"method,omitempty" xml:"method"`

	// 导入失败的错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`
}

func (o Failure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Failure struct{}"
	}

	return strings.Join([]string{"Failure", string(data)}, " ")
}
