package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSyncIamUserRequest Request Object
type UpdateSyncIamUserRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *UpdateSyncRequest `json:"body,omitempty"`
}

func (o UpdateSyncIamUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSyncIamUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateSyncIamUserRequest", string(data)}, " ")
}
