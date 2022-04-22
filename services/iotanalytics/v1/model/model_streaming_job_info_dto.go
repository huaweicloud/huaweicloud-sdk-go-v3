package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实时分析作业基础信息，包括：实时分析作业ID、实时分析作业名称、作业类型等。
type StreamingJobInfoDto struct {

	// 作业ID
	JobId *string `json:"job_id,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 接收数据类型
	JobInputType *string `json:"job_input_type,omitempty"`

	// 作业描述
	JobDescription *string `json:"job_description,omitempty"`

	// 作业状态
	JobState *string `json:"job_state,omitempty"`

	// 操作状态
	Status *string `json:"status,omitempty"`

	// 运行作业的RTU个数
	Rtu *int32 `json:"rtu,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 已停止作业是否有历史缓存数据
	HasSavepoint *bool `json:"has_savepoint,omitempty"`
}

func (o StreamingJobInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamingJobInfoDto struct{}"
	}

	return strings.Join([]string{"StreamingJobInfoDto", string(data)}, " ")
}
