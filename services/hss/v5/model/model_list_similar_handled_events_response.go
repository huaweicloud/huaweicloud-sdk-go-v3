package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSimilarHandledEventsResponse Response Object
type ListSimilarHandledEventsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 相似已处置的告警列表
	DataList       *[]SimilarHandledEvent `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListSimilarHandledEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimilarHandledEventsResponse struct{}"
	}

	return strings.Join([]string{"ListSimilarHandledEventsResponse", string(data)}, " ")
}
