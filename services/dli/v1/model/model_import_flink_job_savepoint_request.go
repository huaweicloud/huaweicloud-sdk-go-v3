package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportFlinkJobSavepointRequest Request Object
type ImportFlinkJobSavepointRequest struct {

	// 非运行阶段的Flink作业的作业ID
	JobId string `json:"job_id"`

	Body *ImportFlinkJobSavepointRequestBody `json:"body,omitempty"`
}

func (o ImportFlinkJobSavepointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportFlinkJobSavepointRequest struct{}"
	}

	return strings.Join([]string{"ImportFlinkJobSavepointRequest", string(data)}, " ")
}
