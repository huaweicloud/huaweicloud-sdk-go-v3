package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListK8sStatefulSetsResponse Response Object
type ListK8sStatefulSetsResponse struct {

	// 有状态工作负载总量
	TotalNum *int64 `json:"total_num,omitempty"`

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 有状态工作负载基本信息列表
	StatefulsetInfoList *[]ServerlessStatefulSetInfo `json:"statefulset_info_list,omitempty"`
	HttpStatusCode      int                          `json:"-"`
}

func (o ListK8sStatefulSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListK8sStatefulSetsResponse struct{}"
	}

	return strings.Join([]string{"ListK8sStatefulSetsResponse", string(data)}, " ")
}
