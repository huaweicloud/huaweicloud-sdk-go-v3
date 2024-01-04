package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSyncIamUserRequest Request Object
type ShowSyncIamUserRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ShowSyncIamUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSyncIamUserRequest struct{}"
	}

	return strings.Join([]string{"ShowSyncIamUserRequest", string(data)}, " ")
}
