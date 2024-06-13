package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StreamJobStatus 流作业的job模型。
type StreamJobStatus struct {

	// 流作业Id。
	StreamId *int64 `json:"stream_id,omitempty"`

	// 流作业状态名称。
	StatusName *string `json:"status_name,omitempty"`

	// 当前状态描述，包含异常状态原因及建议。
	StatusDesc *string `json:"status_desc,omitempty"`
}

func (o StreamJobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamJobStatus struct{}"
	}

	return strings.Join([]string{"StreamJobStatus", string(data)}, " ")
}
