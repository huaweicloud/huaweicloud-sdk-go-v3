package model

import (
	"encoding/json"

	"strings"
)

// 保护实例变更规格请求体
type ResizeProtectedInstanceRequestBody struct {
	Resize *ResizeProtectedInstanceRequestParams `json:"resize"`
}

func (o ResizeProtectedInstanceRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResizeProtectedInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeProtectedInstanceRequestBody", string(data)}, " ")
}
