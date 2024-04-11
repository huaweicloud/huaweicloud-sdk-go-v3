package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataCompareOverviewResponse Response Object
type ListDataCompareOverviewResponse struct {

	// 对比信息数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 信息列表
	DataCompareOverviewInfos *[]DataCompareOverviewInfo `json:"data_compare_overview_infos,omitempty"`
	HttpStatusCode           int                        `json:"-"`
}

func (o ListDataCompareOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataCompareOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListDataCompareOverviewResponse", string(data)}, " ")
}
