package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHoneypotEventsResponse Response Object
type ListHoneypotEventsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 蜜罐事件列表 **取值范围**: 取值0-200个HoneypotEventResponseInfo对象
	DataList       *[]HoneypotEventResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListHoneypotEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHoneypotEventsResponse struct{}"
	}

	return strings.Join([]string{"ListHoneypotEventsResponse", string(data)}, " ")
}
