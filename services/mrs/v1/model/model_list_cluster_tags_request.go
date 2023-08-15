package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterTagsRequest Request Object
type ListClusterTagsRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterTagsRequest", string(data)}, " ")
}
