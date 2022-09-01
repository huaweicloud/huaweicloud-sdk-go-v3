package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListComputingResourcesRequest struct {

	// 计算资源名称。
	ComputingResourceName *string `json:"computing_resource_name,omitempty" xml:"computing_resource_name"`

	// 当前偏移量，默认为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的最大作业个数，范围: [1, 100]。默认值：10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListComputingResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListComputingResourcesRequest", string(data)}, " ")
}
