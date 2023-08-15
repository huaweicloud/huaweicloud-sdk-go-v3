package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTagRequest Request Object
type BatchDeleteResourceTagRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`

	Body *BatchDeleteResourceTags `json:"body,omitempty"`
}

func (o BatchDeleteResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTagRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTagRequest", string(data)}, " ")
}
