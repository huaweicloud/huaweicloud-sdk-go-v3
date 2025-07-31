package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncMultiCloudClusterStatusResponse Response Object
type SyncMultiCloudClusterStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SyncMultiCloudClusterStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncMultiCloudClusterStatusResponse struct{}"
	}

	return strings.Join([]string{"SyncMultiCloudClusterStatusResponse", string(data)}, " ")
}
