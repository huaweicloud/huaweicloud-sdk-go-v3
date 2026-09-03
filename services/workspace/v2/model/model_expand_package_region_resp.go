package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandPackageRegionResp 扩展技能包区域响应。
type ExpandPackageRegionResp struct {

	// 请求总数。
	Total *int32 `json:"total,omitempty"`

	// 成功数量。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails *[]RegionFailedDetail `json:"failed_details,omitempty"`
}

func (o ExpandPackageRegionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandPackageRegionResp struct{}"
	}

	return strings.Join([]string{"ExpandPackageRegionResp", string(data)}, " ")
}
