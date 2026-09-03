package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFtMetricsResponse Response Object
type ShowFtMetricsResponse struct {

	// 训练loss信息
	Loss *interface{} `json:"loss,omitempty"`

	// 评测loss信息
	EvalLoss *interface{} `json:"eval_loss,omitempty"`

	// 训练预估时长信息
	TrainingInfo *interface{} `json:"training_info,omitempty"`

	// 训练进度信息
	TrainProcess *float64 `json:"train_process,omitempty"`

	Data           *FtMetricData `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowFtMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFtMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowFtMetricsResponse", string(data)}, " ")
}
