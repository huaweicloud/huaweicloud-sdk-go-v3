package model

import (
	"encoding/json"

	"strings"
)

type JobDetail struct {
	// 任务ID

	Id string `json:"id"`
	// 任务名称。

	Name string `json:"name"`
	// 任务创建时间，格式为yyyy-mm-ddThh:mm:ssZ。

	Created string `json:"created"`
	// 任务结束时间，格式为yyyy-mm-ddThh:mm:ssZ。

	Ended string `json:"ended"`
	// 任务执行进度。

	Progress string `json:"progress"`

	Instance *JobInstanceInfo `json:"instance"`
}

func (o JobDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobDetail struct{}"
	}

	return strings.Join([]string{"JobDetail", string(data)}, " ")
}
