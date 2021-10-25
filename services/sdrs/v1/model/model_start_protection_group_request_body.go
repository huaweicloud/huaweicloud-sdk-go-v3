package model

import (
	"encoding/json"

	"strings"
)

// 开启保护/重保护请求体
type StartProtectionGroupRequestBody struct {
	// 标识保护组开始保护操作。目前该参数为空。

	StartServerGroup *interface{} `json:"start-server-group"`
}

func (o StartProtectionGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"StartProtectionGroupRequestBody", string(data)}, " ")
}
