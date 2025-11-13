package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSyncRequirementsRequest Request Object
type ListSyncRequirementsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListSyncRequirementsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSyncRequirementsRequest struct{}"
	}

	return strings.Join([]string{"ListSyncRequirementsRequest", string(data)}, " ")
}
