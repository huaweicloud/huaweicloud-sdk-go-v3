package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateLogContents3Response struct {
	// 查询结构化日志结果信息。此处仅为示例，具体参数名称取决于查询的字段。

	Context        *string `json:"context,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLogContents3Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateLogContents3Response struct{}"
	}

	return strings.Join([]string{"UpdateLogContents3Response", string(data)}, " ")
}
