package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBusinessAssetsStatisticResponse Response Object
type ShowBusinessAssetsStatisticResponse struct {

	// 主题域分组的总数
	Count *int32 `json:"count,omitempty"`

	// 主题域分组的统计信息
	SubjectAreaGroupStatistics *[]L1Statistic `json:"subject_area_group_statistics,omitempty"`
	HttpStatusCode             int            `json:"-"`
}

func (o ShowBusinessAssetsStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBusinessAssetsStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowBusinessAssetsStatisticResponse", string(data)}, " ")
}
