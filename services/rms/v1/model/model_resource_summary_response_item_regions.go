package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 区域概要
type ResourceSummaryResponseItemRegions struct {

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 该资源类型在当前区域的数量
	Count *int64 `json:"count,omitempty"`
}

func (o ResourceSummaryResponseItemRegions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSummaryResponseItemRegions struct{}"
	}

	return strings.Join([]string{"ResourceSummaryResponseItemRegions", string(data)}, " ")
}
