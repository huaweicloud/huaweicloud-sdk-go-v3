package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowClusterTagRequest struct {

	// 指定待查询的集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterTagRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterTagRequest", string(data)}, " ")
}
