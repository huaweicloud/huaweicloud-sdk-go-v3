package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSinkTaskRequest struct {
	// 实例转储ID。 请参考[实例生命周期][查询实例]接口返回的数据。

	ConnectorId string `json:"connector_id"`
	// 转储任务ID。

	TaskId string `json:"task_id"`
}

func (o DeleteSinkTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSinkTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteSinkTaskRequest", string(data)}, " ")
}
