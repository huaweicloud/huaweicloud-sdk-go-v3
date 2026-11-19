package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiOverviewStatisticsResponse Response Object
type ShowAiOverviewStatisticsResponse struct {

	// **参数解释**: 按产品类型分组的数据列表 **取值范围**: 不涉及
	DataList       *[]AiOverviewDataList `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowAiOverviewStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiOverviewStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowAiOverviewStatisticsResponse", string(data)}, " ")
}
