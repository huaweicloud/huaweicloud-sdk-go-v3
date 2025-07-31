package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncClusterListResponse Response Object
type SyncClusterListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SyncClusterListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncClusterListResponse struct{}"
	}

	return strings.Join([]string{"SyncClusterListResponse", string(data)}, " ")
}
