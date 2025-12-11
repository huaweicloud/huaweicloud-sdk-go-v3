package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceClustersResponse Response Object
type ShowResourceClustersResponse struct {

	// 已开通的资源列表。
	DedicatedCluster *[]DedicatedCluster `json:"dedicated_cluster,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ShowResourceClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceClustersResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceClustersResponse", string(data)}, " ")
}
