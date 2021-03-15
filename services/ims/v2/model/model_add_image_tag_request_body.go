package model

import (
	"encoding/json"

	"strings"
)

// 请求参数
type AddImageTagRequestBody struct {
	Tag *ResourceTag `json:"tag"`
}

func (o AddImageTagRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddImageTagRequestBody struct{}"
	}

	return strings.Join([]string{"AddImageTagRequestBody", string(data)}, " ")
}
