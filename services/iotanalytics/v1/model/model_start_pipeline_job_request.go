package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartPipelineJobRequest struct {

	// 管道ID
	PipelineId string `json:"pipeline_id"`

	// 运行管道的并发度
	Parallel *int32 `json:"parallel,omitempty"`

	// 运行管道的RTU个数
	Rtu *int32 `json:"rtu,omitempty"`

	// 运行管道作业使用历史缓存数据
	ResumeSavepoint *bool `json:"resume_savepoint,omitempty"`
}

func (o StartPipelineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineJobRequest struct{}"
	}

	return strings.Join([]string{"StartPipelineJobRequest", string(data)}, " ")
}
