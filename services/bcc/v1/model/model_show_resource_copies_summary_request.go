package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceCopiesSummaryRequest Request Object
type ShowResourceCopiesSummaryRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 副本类型
	CopyType *string `json:"copy_type,omitempty"`

	// 统计起始时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime *string `json:"start_time,omitempty"`

	// 统计截止时间。格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ShowResourceCopiesSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceCopiesSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceCopiesSummaryRequest", string(data)}, " ")
}
