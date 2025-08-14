package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityEventsResponse Response Object
type ListSecurityEventsResponse struct {

	// **参数解释**： 告警事件总数 **取值范围**： 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// 事件列表详情
	DataList       *[]EventManagementResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListSecurityEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityEventsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityEventsResponse", string(data)}, " ")
}
