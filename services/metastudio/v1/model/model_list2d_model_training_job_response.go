package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// List2dModelTrainingJobResponse Response Object
type List2dModelTrainingJobResponse struct {

	// 分身数字人模型训练任务数量。
	Count *int32 `json:"count,omitempty"`

	// 分身数字人模型训练任务列表。
	Jobs *[]TrainingJobBasicInfo `json:"jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o List2dModelTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "List2dModelTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"List2dModelTrainingJobResponse", string(data)}, " ")
}
