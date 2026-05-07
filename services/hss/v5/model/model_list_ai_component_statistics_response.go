package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiComponentStatisticsResponse Response Object
type ListAiComponentStatisticsResponse struct {

	// **参数解释**： AI组件列表总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： AI组件列表
	DataList       *[]AiStatisticInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListAiComponentStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiComponentStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListAiComponentStatisticsResponse", string(data)}, " ")
}
