package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionsRequest Request Object
type ListActionsRequest struct {

	// 指定查询集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定查询起始值，默认值为0。
	Start *string `json:"start,omitempty"`

	// 指定查询个数，默认值为10。
	Limit *string `json:"limit,omitempty"`
}

func (o ListActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsRequest struct{}"
	}

	return strings.Join([]string{"ListActionsRequest", string(data)}, " ")
}
