package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskTypeSummaryResponseBodyFilter 过滤条件
type ShowTaskTypeSummaryResponseBodyFilter struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	ResourceType *ResourceTypeEnum `json:"resource_type,omitempty"`

	// 起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowTaskTypeSummaryResponseBodyFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskTypeSummaryResponseBodyFilter struct{}"
	}

	return strings.Join([]string{"ShowTaskTypeSummaryResponseBodyFilter", string(data)}, " ")
}
