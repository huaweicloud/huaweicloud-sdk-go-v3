package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业的输出配置，至少需要配置一项输出方式。
type TaskOutput struct {
	Obs *TaskOutputObs `json:"obs,omitempty" xml:"obs"`

	Dis *TaskOutputDis `json:"dis,omitempty" xml:"dis"`

	Webhook *TaskOutputWebhook `json:"webhook,omitempty" xml:"webhook"`

	Hosting *TaskOutputHosting `json:"hosting,omitempty" xml:"hosting"`

	Localpath *TaskOutputLocalpath `json:"localpath,omitempty" xml:"localpath"`
}

func (o TaskOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutput struct{}"
	}

	return strings.Join([]string{"TaskOutput", string(data)}, " ")
}
