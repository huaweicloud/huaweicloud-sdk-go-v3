package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCssClusterResponse Response Object
type ListCssClusterResponse struct {

	// CSS集群列表
	CssClusters *[]CssClusterDto `json:"css_clusters,omitempty"`

	// css集群总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCssClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCssClusterResponse struct{}"
	}

	return strings.Join([]string{"ListCssClusterResponse", string(data)}, " ")
}
