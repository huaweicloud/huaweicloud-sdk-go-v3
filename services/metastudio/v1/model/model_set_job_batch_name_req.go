package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetJobBatchNameReq 设置任务批次信息
type SetJobBatchNameReq struct {

	// 批次名称
	BatchName *string `json:"batch_name,omitempty"`

	// 任务id列表
	JobIds *[]string `json:"job_ids,omitempty"`
}

func (o SetJobBatchNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetJobBatchNameReq struct{}"
	}

	return strings.Join([]string{"SetJobBatchNameReq", string(data)}, " ")
}
