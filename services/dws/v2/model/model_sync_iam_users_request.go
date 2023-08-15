package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncIamUsersRequest Request Object
type SyncIamUsersRequest struct {

	// cluster_id
	ClusterId string `json:"cluster_id"`
}

func (o SyncIamUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncIamUsersRequest struct{}"
	}

	return strings.Join([]string{"SyncIamUsersRequest", string(data)}, " ")
}
