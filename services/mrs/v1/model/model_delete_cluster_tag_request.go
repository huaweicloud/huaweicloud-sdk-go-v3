package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteClusterTagRequest struct {
	// 集群ID。

	ClusterId string `json:"cluster_id"`
	// 键。标签的key值

	Key string `json:"key"`
}

func (o DeleteClusterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterTagRequest", string(data)}, " ")
}
