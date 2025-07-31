package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMultiCloudClustersResponse Response Object
type CreateMultiCloudClustersResponse struct {

	// 集群id
	ClusterId      *string `json:"cluster_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMultiCloudClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMultiCloudClustersResponse struct{}"
	}

	return strings.Join([]string{"CreateMultiCloudClustersResponse", string(data)}, " ")
}
