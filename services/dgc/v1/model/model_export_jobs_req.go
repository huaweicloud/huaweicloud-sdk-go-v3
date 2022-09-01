package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportJobsReq struct {
	JobList *[]string `json:"jobList,omitempty" xml:"jobList"`

	// 是否导出作业依赖的脚本和资源
	ExportDepend *bool `json:"exportDepend,omitempty" xml:"exportDepend"`
}

func (o ExportJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobsReq struct{}"
	}

	return strings.Join([]string{"ExportJobsReq", string(data)}, " ")
}
