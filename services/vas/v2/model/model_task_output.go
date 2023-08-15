package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskOutput 作业的输出配置，至少需要配置一项输出方式。
type TaskOutput struct {
	Obs *TaskOutputObs `json:"obs,omitempty"`

	Dis *TaskOutputDis `json:"dis,omitempty"`

	Webhook *TaskOutputWebhook `json:"webhook,omitempty"`

	Hosting *TaskOutputHosting `json:"hosting,omitempty"`

	Localpath *TaskOutputLocalpath `json:"localpath,omitempty"`
}

func (o TaskOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutput struct{}"
	}

	return strings.Join([]string{"TaskOutput", string(data)}, " ")
}
