package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKubernetesServiceDetailsResponse Response Object
type ListKubernetesServiceDetailsResponse struct {

	// 服务总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 服务列表
	ServiceInfoList *[]KubernetesServiceInfo `json:"service_info_list,omitempty"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ListKubernetesServiceDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKubernetesServiceDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListKubernetesServiceDetailsResponse", string(data)}, " ")
}
