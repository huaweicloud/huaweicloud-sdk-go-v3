package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudTableClusterResponse Response Object
type CreateCloudTableClusterResponse struct {

	// 集群ID
	ClusterId      *string `json:"cluster_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCloudTableClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudTableClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudTableClusterResponse", string(data)}, " ")
}
