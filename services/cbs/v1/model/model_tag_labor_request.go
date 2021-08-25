package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type TagLaborRequest struct {
	// qabot编号，UUID格式。

	QabotId string `json:"qabot_id"`
	// 请求ID，由问答机器人会话生成。

	RequestId string `json:"request_id"`
}

func (o TagLaborRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TagLaborRequest struct{}"
	}

	return strings.Join([]string{"TagLaborRequest", string(data)}, " ")
}
