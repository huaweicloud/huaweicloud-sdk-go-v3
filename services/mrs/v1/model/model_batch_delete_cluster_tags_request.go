package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteClusterTagsRequest struct {
	// 集群ID。获取方法，请参见[获取集群ID](mrs_02_9001.xml)。

	ClusterId string `json:"cluster_id"`

	Body *BatchDeleteClusterTagsReq `json:"body,omitempty"`
}

func (o BatchDeleteClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterTagsRequest", string(data)}, " ")
}
