package model

import (
	"encoding/json"

	"strings"
)

type BatchTagActionAddRequestBody struct {
	// 操作标识（区分大小写）：创建时为“create”。
	Action string `json:"action"`
	// 标签列表。单个实例总标签数上限10个。
	Tags []TagWithKeyValue `json:"tags"`
}

func (o BatchTagActionAddRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchTagActionAddRequestBody struct{}"
	}

	return strings.Join([]string{"BatchTagActionAddRequestBody", string(data)}, " ")
}
