package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateClusterTagsRequest Request Object
type BatchCreateClusterTagsRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`

	Body *BatchCreateClusterTagsReq `json:"body,omitempty"`
}

func (o BatchCreateClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterTagsRequest", string(data)}, " ")
}
