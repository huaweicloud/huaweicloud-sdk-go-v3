package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelinesRequest Request Object
type ListPipelinesRequest struct {

	// 指定查询集群ID。
	ClusterId string `json:"cluster_id"`

	// 指定查询起始值，默认值为0。
	Offset *string `json:"offset,omitempty"`

	// 指定查询个数，默认值为10。
	Limit *string `json:"limit,omitempty"`
}

func (o ListPipelinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesRequest struct{}"
	}

	return strings.Join([]string{"ListPipelinesRequest", string(data)}, " ")
}
