package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncStatusRequest Request Object
type ListSyncStatusRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListSyncStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncStatusRequest struct{}"
	}

	return strings.Join([]string{"ListSyncStatusRequest", string(data)}, " ")
}
