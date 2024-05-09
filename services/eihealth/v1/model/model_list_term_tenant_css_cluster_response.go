package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTermTenantCssClusterResponse Response Object
type ListTermTenantCssClusterResponse struct {

	// 最终租户CSS集群列表
	CssClusters *[]TermTenantCssClusterDto `json:"css_clusters,omitempty"`

	// 最终租户css集群总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTermTenantCssClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTermTenantCssClusterResponse struct{}"
	}

	return strings.Join([]string{"ListTermTenantCssClusterResponse", string(data)}, " ")
}
