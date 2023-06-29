package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvInstancesResponse Response Object
type ListEnvInstancesResponse struct {

	// 实例信息列表。
	InstanceInfoList *[]InstanceInfo `json:"instance_info_list,omitempty"`

	// 实例总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 在线实例总数。
	OnlineCount *int32 `json:"online_count,omitempty"`

	// 离线实例总数。
	OfflineCount *int32 `json:"offline_count,omitempty"`

	// 停止实例总数。
	DisableCount   *int32 `json:"disable_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnvInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListEnvInstancesResponse", string(data)}, " ")
}
