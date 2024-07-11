package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostClustersResponse Response Object
type ListHostClustersResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 主机集群个数
	Total *int32 `json:"total,omitempty"`

	// 主机集群详情响应体
	Result         *[]HostClusterInfo `json:"result,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListHostClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostClustersResponse struct{}"
	}

	return strings.Join([]string{"ListHostClustersResponse", string(data)}, " ")
}
