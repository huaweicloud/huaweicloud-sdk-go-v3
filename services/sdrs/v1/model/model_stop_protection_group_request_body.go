package model

import (
	"encoding/json"

	"strings"
)

// 保护组停止保护请求体
type StopProtectionGroupRequestBody struct {
	// 标识保护组停止保护操作。该参数目前默认值为空。

	StopServerGroup *interface{} `json:"stop-server-group"`
}

func (o StopProtectionGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"StopProtectionGroupRequestBody", string(data)}, " ")
}
