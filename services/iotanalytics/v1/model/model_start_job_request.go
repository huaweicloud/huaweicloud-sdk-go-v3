package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartJobRequest struct {
	// 作业ID

	JobId string `json:"job_id"`
	// 运行作业的并发度

	Parallel *int32 `json:"parallel,omitempty"`
	// 运行作业的RTU个数

	Rtu *int32 `json:"rtu,omitempty"`
	// 运行作业使用历史缓存数据

	ResumeSavepoint *bool `json:"resume_savepoint,omitempty"`
}

func (o StartJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobRequest struct{}"
	}

	return strings.Join([]string{"StartJobRequest", string(data)}, " ")
}
