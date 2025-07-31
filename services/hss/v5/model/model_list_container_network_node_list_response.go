package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerNetworkNodeListResponse Response Object
type ListContainerNetworkNodeListResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 数据最近同步时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 网络节点列表
	DataList       *[]NetworkNodeInfo `json:"data_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListContainerNetworkNodeListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerNetworkNodeListResponse struct{}"
	}

	return strings.Join([]string{"ListContainerNetworkNodeListResponse", string(data)}, " ")
}
