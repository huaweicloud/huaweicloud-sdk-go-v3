package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllFlowsRequest Request Object
type ShowAllFlowsRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset int32 `json:"offset"`

	// 每页显示的条目数量，limit大于等于1
	Limit int32 `json:"limit"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 流名称，支持模糊查询
	Name *string `json:"name,omitempty"`

	// 是否包含子流程
	HaveChildFlow *bool `json:"have_child_flow,omitempty"`

	// ids
	Ids *string `json:"ids,omitempty"`
}

func (o ShowAllFlowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllFlowsRequest struct{}"
	}

	return strings.Join([]string{"ShowAllFlowsRequest", string(data)}, " ")
}
