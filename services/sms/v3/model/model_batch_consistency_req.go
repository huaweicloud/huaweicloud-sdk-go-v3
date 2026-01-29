package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchConsistencyReq 需要获取一致性校验结果任务的请求参数
type BatchConsistencyReq struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 源端ID
	SourceId *string `json:"source_id,omitempty"`

	// 源端名称
	SourceName *string `json:"source_name,omitempty"`
}

func (o BatchConsistencyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchConsistencyReq struct{}"
	}

	return strings.Join([]string{"BatchConsistencyReq", string(data)}, " ")
}
