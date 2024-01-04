package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVoiceTrainingJobResponse Response Object
type ListVoiceTrainingJobResponse struct {

	// 满足查询要求的任务总数。
	Count *int32 `json:"count,omitempty"`

	// 分身数字人模型训练任务列表。
	Jobs           *[]TrainingJobInfo `json:"jobs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListVoiceTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVoiceTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"ListVoiceTrainingJobResponse", string(data)}, " ")
}
