package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实时分析作业基础信息，包括：实时分析作业ID、实时分析作业名称、作业类型等。
type StreamingJobInfoDto struct {

	// 作业ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 作业名称
	JobName *string `json:"job_name,omitempty" xml:"job_name"`

	// 接收数据类型
	JobInputType *string `json:"job_input_type,omitempty" xml:"job_input_type"`

	// 作业描述
	JobDescription *string `json:"job_description,omitempty" xml:"job_description"`

	// 作业状态
	JobState *string `json:"job_state,omitempty" xml:"job_state"`

	// 操作状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 运行作业的RTU个数
	Rtu *int32 `json:"rtu,omitempty" xml:"rtu"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty" xml:"modified_time"`

	// 用户ID
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 已停止作业是否有历史缓存数据
	HasSavepoint *bool `json:"has_savepoint,omitempty" xml:"has_savepoint"`
}

func (o StreamingJobInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StreamingJobInfoDto struct{}"
	}

	return strings.Join([]string{"StreamingJobInfoDto", string(data)}, " ")
}
