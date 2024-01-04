package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSyncIamUserRequest Request Object
type CancelSyncIamUserRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *CancelSyncRequest `json:"body,omitempty"`
}

func (o CancelSyncIamUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSyncIamUserRequest struct{}"
	}

	return strings.Join([]string{"CancelSyncIamUserRequest", string(data)}, " ")
}
