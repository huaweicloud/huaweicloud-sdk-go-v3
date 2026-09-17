package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNamespacesResponse Response Object
type ListClusterNamespacesResponse struct {

	// 集群命名空间列表
	Namespaces     *[]QueryNamespaceResp `json:"namespaces,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListClusterNamespacesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNamespacesResponse struct{}"
	}

	return strings.Join([]string{"ListClusterNamespacesResponse", string(data)}, " ")
}
