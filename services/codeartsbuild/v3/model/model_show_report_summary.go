package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportSummary 单元测试报告
type ShowReportSummary struct {

	// 任务编号
	JobId *string `json:"job_id,omitempty"`

	// 构建编号
	BuildNo *int32 `json:"build_no,omitempty"`

	// 步骤名称，对应构建步骤，例如stage2
	StageName *string `json:"stage_name,omitempty"`

	// 报告编号
	RootId *string `json:"root_id,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 成功数量
	Success *int32 `json:"success,omitempty"`

	// 失败数量
	Failures *int32 `json:"failures,omitempty"`

	// 错误数量
	Errors *int32 `json:"errors,omitempty"`

	// 其他
	Others *int32 `json:"others,omitempty"`

	// 执行耗时
	ExecutionTime *int32 `json:"execution_time,omitempty"`

	// 生成时间
	GenrateTime *string `json:"genrate_time,omitempty"`

	// 局点
	Region *string `json:"region,omitempty"`

	// 是否通过
	IsSuccess *bool `json:"is_success,omitempty"`
}

func (o ShowReportSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSummary struct{}"
	}

	return strings.Join([]string{"ShowReportSummary", string(data)}, " ")
}
