package model

import (
	"encoding/json"

	"strings"
)

type ExportJobsReq struct {
	JobList *[]string `json:"jobList,omitempty"`
	// 是否导出作业依赖的脚本和资源

	ExportDepend *bool `json:"exportDepend,omitempty"`
}

func (o ExportJobsReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportJobsReq struct{}"
	}

	return strings.Join([]string{"ExportJobsReq", string(data)}, " ")
}
