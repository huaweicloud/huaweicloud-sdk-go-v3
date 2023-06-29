package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersRequest Request Object
type ListClustersRequest struct {

	// 边缘集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 边缘集群状态
	State *string `json:"state,omitempty"`

	// 每页记录数，默认值为10，取值区间为1-1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
