package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterFlavorRequest Request Object
type ShowClusterFlavorRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o ShowClusterFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterFlavorRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterFlavorRequest", string(data)}, " ")
}
