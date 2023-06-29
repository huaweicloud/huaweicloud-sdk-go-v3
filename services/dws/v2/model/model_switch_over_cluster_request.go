package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOverClusterRequest Request Object
type SwitchOverClusterRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`
}

func (o SwitchOverClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOverClusterRequest struct{}"
	}

	return strings.Join([]string{"SwitchOverClusterRequest", string(data)}, " ")
}
