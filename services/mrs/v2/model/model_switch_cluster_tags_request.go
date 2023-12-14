package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchClusterTagsRequest Request Object
type SwitchClusterTagsRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	Body *ModifyDefaultTagsRequestBody `json:"body,omitempty"`
}

func (o SwitchClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"SwitchClusterTagsRequest", string(data)}, " ")
}
