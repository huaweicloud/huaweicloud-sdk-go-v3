package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSummaryResponseItem 资源概要信息
type ResourceSummaryResponseItem struct {

	// 云服务名称
	Provider *string `json:"provider,omitempty"`

	// 资源类型列表
	Types *[]ResourceSummaryResponseItemTypes `json:"types,omitempty"`
}

func (o ResourceSummaryResponseItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSummaryResponseItem struct{}"
	}

	return strings.Join([]string{"ResourceSummaryResponseItem", string(data)}, " ")
}
