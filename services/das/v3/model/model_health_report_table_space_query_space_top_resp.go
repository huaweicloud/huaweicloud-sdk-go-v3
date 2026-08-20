package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HealthReportTableSpaceQuerySpaceTopResp struct {

	// 库/表大小Top列表。
	TopDataList *[]HealthReportTableSpaceTopDataDto `json:"top_data_list,omitempty"`

	// 采集时间
	CollectTimestamp *int64 `json:"collect_timestamp,omitempty"`

	// 总大小。
	TotalCount *int64 `json:"total_count,omitempty"`
}

func (o HealthReportTableSpaceQuerySpaceTopResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthReportTableSpaceQuerySpaceTopResp struct{}"
	}

	return strings.Join([]string{"HealthReportTableSpaceQuerySpaceTopResp", string(data)}, " ")
}
