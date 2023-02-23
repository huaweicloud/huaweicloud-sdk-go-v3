package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源类型
type ResourceSummaryResponseItemTypes struct {

	// 资源类型名称
	Type *string `json:"type,omitempty"`

	// 区域列表
	Regions *[]ResourceSummaryResponseItemRegions `json:"regions,omitempty"`
}

func (o ResourceSummaryResponseItemTypes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSummaryResponseItemTypes struct{}"
	}

	return strings.Join([]string{"ResourceSummaryResponseItemTypes", string(data)}, " ")
}
