package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceInternalEndpointsRequest Request Object
type ListInstanceInternalEndpointsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 起始索引，默认值为0。**注意：offset和limit参数需要配套使用。**
	Offset *int32 `json:"offset,omitempty"`

	// 返回条数，默认为100，最大值为100。**注意：offset和limit参数需要配套使用。**
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListInstanceInternalEndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceInternalEndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceInternalEndpointsRequest", string(data)}, " ")
}
