package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSinkTaskDetailRequest struct {
	// 实例转储ID。 请参考[实例生命周期][查询实例]接口返回的数据。

	ConnectorId string `json:"connector_id"`
	// 转储任务ID。

	TaskId string `json:"task_id"`
}

func (o ShowSinkTaskDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRequest", string(data)}, " ")
}
