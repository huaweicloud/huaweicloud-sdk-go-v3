package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateResourceTagRequest Request Object
type BatchCreateResourceTagRequest struct {

	// 集群的ID
	ClusterId string `json:"cluster_id"`

	Body *BatchCreateResourceTags `json:"body,omitempty"`
}

func (o BatchCreateResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagRequest", string(data)}, " ")
}
