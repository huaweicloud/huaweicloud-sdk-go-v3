package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubTrainingJobInfoDto 子任务信息
type SubTrainingJobInfoDto struct {

	// 任务id
	JobId *string `json:"job_id,omitempty"`

	// 子任务类型
	SubJobType *string `json:"sub_job_type,omitempty"`

	// 子任务心跳id
	HeartBeatId *string `json:"heart_beat_id,omitempty"`

	// 任务状态
	State *string `json:"state,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最后更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`
}

func (o SubTrainingJobInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTrainingJobInfoDto struct{}"
	}

	return strings.Join([]string{"SubTrainingJobInfoDto", string(data)}, " ")
}
