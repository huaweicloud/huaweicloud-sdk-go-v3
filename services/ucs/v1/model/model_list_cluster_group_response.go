package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterGroupResponse Response Object
type ListClusterGroupResponse struct {

	// 容器舰队列表信息
	Items *[]ClusterGroup `json:"items,omitempty"`

	// 所有页的结果的总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClusterGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterGroupResponse struct{}"
	}

	return strings.Join([]string{"ListClusterGroupResponse", string(data)}, " ")
}
