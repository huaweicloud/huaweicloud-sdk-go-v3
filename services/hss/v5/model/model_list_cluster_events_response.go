package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterEventsResponse Response Object
type ListClusterEventsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 集群安全事件列表
	DataList       *[]ClusterEventResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListClusterEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterEventsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterEventsResponse", string(data)}, " ")
}
