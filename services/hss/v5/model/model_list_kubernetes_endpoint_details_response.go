package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKubernetesEndpointDetailsResponse Response Object
type ListKubernetesEndpointDetailsResponse struct {

	// 端点总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 端点列表
	EndpointInfoList *[]KubernetesEndpointInfo `json:"endpoint_info_list,omitempty"`
	HttpStatusCode   int                       `json:"-"`
}

func (o ListKubernetesEndpointDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKubernetesEndpointDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListKubernetesEndpointDetailsResponse", string(data)}, " ")
}
