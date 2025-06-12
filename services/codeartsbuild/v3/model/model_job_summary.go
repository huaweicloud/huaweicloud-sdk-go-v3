package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobSummary 构建成功率
type JobSummary struct {

	// 构建成功率
	AvgSuccessRatio *int32 `json:"avg_success_ratio,omitempty"`

	// 构建总时长
	BuildNo *int32 `json:"build_no,omitempty"`

	// 任务总数
	JobTotal *int32 `json:"job_total,omitempty"`

	// 版本
	VersionTotal *string `json:"version_total,omitempty"`
}

func (o JobSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSummary struct{}"
	}

	return strings.Join([]string{"JobSummary", string(data)}, " ")
}
