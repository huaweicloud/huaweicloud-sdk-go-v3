package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListJunitCoverageSummaryResultUnitSummaryList struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 构建任务的构建编号，从1开始，每次构建递增1
	BuildNo *int32 `json:"build_no,omitempty"`

	// stage名称
	StageName *string `json:"stage_name,omitempty"`

	// 资源ID，下载覆盖率报告时使用
	RootId *string `json:"root_id,omitempty"`

	// 租户所在region
	Region *string `json:"region,omitempty"`
}

func (o ListJunitCoverageSummaryResultUnitSummaryList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJunitCoverageSummaryResultUnitSummaryList struct{}"
	}

	return strings.Join([]string{"ListJunitCoverageSummaryResultUnitSummaryList", string(data)}, " ")
}
