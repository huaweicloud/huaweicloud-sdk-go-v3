package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeClusterEventsResponse Response Object
type ChangeClusterEventsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 最近更新时间 **取值范围**: 最小值0，最大值9223372036854775807
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// **参数解释**: 集群安全事件列表 **取值范围**: 取值0-10000个ClusterEventResponseInfo对象
	DataList       *[]ClusterEventResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ChangeClusterEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeClusterEventsResponse struct{}"
	}

	return strings.Join([]string{"ChangeClusterEventsResponse", string(data)}, " ")
}
