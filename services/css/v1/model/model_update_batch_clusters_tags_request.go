package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateBatchClustersTagsRequest struct {

	// 指定待添加的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *BatchAddOrDeleteTagOnClusterReq `json:"body,omitempty"`
}

func (o UpdateBatchClustersTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBatchClustersTagsRequest struct{}"
	}

	return strings.Join([]string{"UpdateBatchClustersTagsRequest", string(data)}, " ")
}
