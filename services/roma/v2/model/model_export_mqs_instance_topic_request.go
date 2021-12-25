package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportMqsInstanceTopicRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 待导出的topic。多个topic以“,”分隔。默认导出所有的topic。

	Name *string `json:"name,omitempty"`
}

func (o ExportMqsInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportMqsInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"ExportMqsInstanceTopicRequest", string(data)}, " ")
}
