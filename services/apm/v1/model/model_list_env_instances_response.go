package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEnvInstancesResponse struct {

	// 实例信息列表
	InstanceInfoList *[]InstanceInfo `json:"instance_info_list,omitempty" xml:"instance_info_list"`

	// 实例总数
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 在线实例总数
	OnlineCount *int32 `json:"online_count,omitempty" xml:"online_count"`

	// 离线实例总数
	OfflineCount *int32 `json:"offline_count,omitempty" xml:"offline_count"`

	// 停止实例总受
	DisableCount   *int32 `json:"disable_count,omitempty" xml:"disable_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnvInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListEnvInstancesResponse", string(data)}, " ")
}
