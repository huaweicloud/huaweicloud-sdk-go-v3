package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateClustersTagsRequest struct {
	// 指定待添加的集群ID。

	ClusterId string `json:"cluster_id"`

	Body *TagReq `json:"body,omitempty"`
}

func (o CreateClustersTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClustersTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateClustersTagsRequest", string(data)}, " ")
}
