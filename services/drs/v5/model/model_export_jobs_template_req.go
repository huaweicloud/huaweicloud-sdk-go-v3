package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportJobsTemplateReq struct {

	// 需要导出的任务ID列表。
	Jobs []string `json:"jobs"`
}

func (o ExportJobsTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportJobsTemplateReq struct{}"
	}

	return strings.Join([]string{"ExportJobsTemplateReq", string(data)}, " ")
}
