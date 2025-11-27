package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterRequest Request Object
type UpdateClusterRequest struct {

	// 集群ID
	Clusterid string `json:"clusterid"`

	Body *UpdateClusterRequestBody `json:"body,omitempty"`
}

func (o UpdateClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterRequest", string(data)}, " ")
}
