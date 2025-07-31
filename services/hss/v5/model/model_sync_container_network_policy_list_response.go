package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncContainerNetworkPolicyListResponse Response Object
type SyncContainerNetworkPolicyListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SyncContainerNetworkPolicyListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncContainerNetworkPolicyListResponse struct{}"
	}

	return strings.Join([]string{"SyncContainerNetworkPolicyListResponse", string(data)}, " ")
}
