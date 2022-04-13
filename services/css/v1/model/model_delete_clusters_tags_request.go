package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteClustersTagsRequest struct {
	// 指定删除集群ID。

	ClusterId string `json:"cluster_id"`
	// 需要删除的标签名。

	Key string `json:"key"`
}

func (o DeleteClustersTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClustersTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteClustersTagsRequest", string(data)}, " ")
}
