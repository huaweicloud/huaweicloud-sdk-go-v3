package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExportFunctionRequest struct {
	// 是否导出函数配置

	Config *bool `json:"config,omitempty"`
	// 是否导出函数代码

	Code *bool `json:"code,omitempty"`
}

func (o ExportFunctionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportFunctionRequest struct{}"
	}

	return strings.Join([]string{"ExportFunctionRequest", string(data)}, " ")
}
