package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelatedEventsResponse Response Object
type ListRelatedEventsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 相关的告警列表 **取值范围**: 最小值0，最大值10
	DataList       *[]RelatedEventInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListRelatedEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedEventsResponse struct{}"
	}

	return strings.Join([]string{"ListRelatedEventsResponse", string(data)}, " ")
}
