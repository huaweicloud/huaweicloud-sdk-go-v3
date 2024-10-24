package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecsRequest Request Object
type ListSpecsRequest struct {

	// 规格编码
	SpecCode *string `json:"spec_code,omitempty"`

	// 通过资源规格类型检索
	SpecTypes *[]string `json:"spec_types,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 规格使用场景，不填表示不限制： COMPUTE: 用于购买Ray计算资源时配置的物理节点规格 ENDPOINT: 用于创建Endpoint时配置的资源组规格大小
	Scenario *SpecScenario `json:"scenario,omitempty"`
}

func (o ListSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecsRequest struct{}"
	}

	return strings.Join([]string{"ListSpecsRequest", string(data)}, " ")
}
