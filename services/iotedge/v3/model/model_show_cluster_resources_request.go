package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterResourcesRequest Request Object
type ShowClusterResourcesRequest struct {

	// 查询资源组
	ResourceGroup *string `json:"resource_group,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认值为10，取值区间为1-100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowClusterResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterResourcesRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterResourcesRequest", string(data)}, " ")
}
