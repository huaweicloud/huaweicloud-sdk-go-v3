package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterListResponse Response Object
type ShowClusterListResponse struct {

	// 集群成员的列表
	Items *[]Cluster `json:"items,omitempty"`

	// 集群总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowClusterListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterListResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterListResponse", string(data)}, " ")
}
