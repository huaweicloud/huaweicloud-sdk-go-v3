package model

import (
	"encoding/json"

	"strings"
)

// 添加保护实例标签请求体
type ProtectedInstanceAddTagsRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o ProtectedInstanceAddTagsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProtectedInstanceAddTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceAddTagsRequestBody", string(data)}, " ")
}
