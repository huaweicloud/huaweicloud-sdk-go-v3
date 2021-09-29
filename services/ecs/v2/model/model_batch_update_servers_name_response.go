package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchUpdateServersNameResponse struct {
	// 提交请求成功后返回的响应列表。

	Response       *[]ServerId `json:"response,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o BatchUpdateServersNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUpdateServersNameResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateServersNameResponse", string(data)}, " ")
}
