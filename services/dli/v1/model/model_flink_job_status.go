package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobStatus 创建作业的job模型。
type FlinkJobStatus struct {

	// 作业Id。
	JobId *int64 `json:"job_id,omitempty"`

	// 作业状态名称。
	StatusName *string `json:"status_name,omitempty"`

	// 当前状态描述，包含异常状态原因及建议。
	StatusDesc *string `json:"status_desc,omitempty"`
}

func (o FlinkJobStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobStatus struct{}"
	}

	return strings.Join([]string{"FlinkJobStatus", string(data)}, " ")
}
