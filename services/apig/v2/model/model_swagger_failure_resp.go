package model

import (
	"encoding/json"

	"strings"
)

type SwaggerFailureResp struct {
	// API请求路径

	Path *string `json:"path,omitempty"`
	// 导入失败的错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
	// API请求方法

	Method *string `json:"method,omitempty"`
	// 导入失败的错误码

	ErrorCode *string `json:"error_code,omitempty"`
}

func (o SwaggerFailureResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SwaggerFailureResp struct{}"
	}

	return strings.Join([]string{"SwaggerFailureResp", string(data)}, " ")
}
