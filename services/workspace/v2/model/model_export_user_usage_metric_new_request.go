package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserUsageMetricNewRequest Request Object
type ExportUserUsageMetricNewRequest struct {

	// 查询起始时间(0时区)。
	StartTime string `json:"start_time"`

	// 查询截至时间(0时区)。
	EndTime string `json:"end_time"`

	// 用户名(模糊匹配)。
	Username *string `json:"username,omitempty"`

	// 使用时长最小值。
	UsageMinHours *int32 `json:"usage_min_hours,omitempty"`

	// 使用时长最大值 usage_min_hours和usage_max_hours同时存在时,usage_max_hours必须大于等于usage_min_hours
	UsageMaxHours *int32 `json:"usage_max_hours,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 按照指标进行排序 * `user_usage` -  按照用户使用时长排序
	SortField *string `json:"sort_field,omitempty"`

	// 按照指标进行排序的方向;需配合sort_field起义使用 * `DESC` - 降序返回数据 * `ASC` -  升序返回数据
	SortType *string `json:"sort_type,omitempty"`
}

func (o ExportUserUsageMetricNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserUsageMetricNewRequest struct{}"
	}

	return strings.Join([]string{"ExportUserUsageMetricNewRequest", string(data)}, " ")
}
