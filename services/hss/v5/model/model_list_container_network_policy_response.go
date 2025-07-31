package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerNetworkPolicyResponse Response Object
type ListContainerNetworkPolicyResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 网络策略列表
	DataList *[]GetNetworkPolicy `json:"data_list,omitempty"`

	// 数据最近同步时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListContainerNetworkPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerNetworkPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListContainerNetworkPolicyResponse", string(data)}, " ")
}
