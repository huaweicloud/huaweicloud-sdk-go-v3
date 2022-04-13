package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListActionsRequest struct {
	// 指定查询集群ID。

	ClusterId string `json:"cluster_id"`
}

func (o ListActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsRequest struct{}"
	}

	return strings.Join([]string{"ListActionsRequest", string(data)}, " ")
}
