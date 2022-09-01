package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业的输出配置展示
type TaskOutputForDisplay struct {
	Obs *TaskOutputObs `json:"obs,omitempty" xml:"obs"`

	Dis *TaskOutputDis `json:"dis,omitempty" xml:"dis"`

	Webhook *TaskOutputWebhook `json:"webhook,omitempty" xml:"webhook"`

	Localpath *TaskOutputLocalpath `json:"localpath,omitempty" xml:"localpath"`
}

func (o TaskOutputForDisplay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputForDisplay struct{}"
	}

	return strings.Join([]string{"TaskOutputForDisplay", string(data)}, " ")
}
