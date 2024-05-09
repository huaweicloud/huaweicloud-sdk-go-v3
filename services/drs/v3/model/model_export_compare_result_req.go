package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCompareResultReq 生成对比任务结果文件请求体。
type ExportCompareResultReq struct {

	// 对比任务类型： - contents： 内容对比。 - lines：行数对比。 - random：抽样对比。 - objects_comparison：对象对比。
	CompareType string `json:"compare_type"`

	// 对比任务的ID，内容对比、抽样对比、行数对比场景必填。
	CompareJobId *string `json:"compare_job_id,omitempty"`

	// 时区，如GMT+08:00，用于生成当前时间标识，拼接到文件名称中。
	TimeZone *string `json:"time_zone,omitempty"`
}

func (o ExportCompareResultReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCompareResultReq struct{}"
	}

	return strings.Join([]string{"ExportCompareResultReq", string(data)}, " ")
}
