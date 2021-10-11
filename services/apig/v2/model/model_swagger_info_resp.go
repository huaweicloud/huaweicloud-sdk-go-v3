package model

import (
	"encoding/json"

	"strings"
)

// swagger文档导入结果  暂不支持
type SwaggerInfoResp struct {
	// swagger文档编号

	Id *string `json:"id,omitempty"`
	// 导入结果说明

	Result *string `json:"result,omitempty"`
}

func (o SwaggerInfoResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SwaggerInfoResp struct{}"
	}

	return strings.Join([]string{"SwaggerInfoResp", string(data)}, " ")
}
