package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskOutputObs struct {

	// 任务运行结果输出的OBS桶名。
	Bucket string `json:"bucket"`

	// 任务运行结果输出的OBS路径。
	Path string `json:"path"`
}

func (o TaskOutputObs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputObs struct{}"
	}

	return strings.Join([]string{"TaskOutputObs", string(data)}, " ")
}
