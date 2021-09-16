package model

import (
	"encoding/json"

	"strings"
)

// 任务类响应通用返回体。
type JobResult struct {
	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。

	JobId string `json:"job_id"`
}

func (o JobResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobResult struct{}"
	}

	return strings.Join([]string{"JobResult", string(data)}, " ")
}
