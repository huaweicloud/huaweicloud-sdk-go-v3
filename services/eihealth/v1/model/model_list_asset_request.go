package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAssetRequest struct {

	// 资产类别，支持IMAGE/APP/WORKFLOW/DATASET，支持查询多个，以','分割
	Categories *string `json:"categories,omitempty"`

	// 关键字，支持在资产名、资产标题、短描述、长描述中搜索
	KeyWord *string `json:"key_word,omitempty"`

	// 标签，支持查询多个，以','分割
	Labels *string `json:"labels,omitempty"`

	// 返回记录限制
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 发布者，支持查询多个，以','分割
	Publishers *string `json:"publishers,omitempty"`

	// 查询范围，支持PUBLIC/INTERNAL
	Scope string `json:"scope"`

	// 供应商，支持查询多个，以','分割
	VendorIds *string `json:"vendor_ids,omitempty"`
}

func (o ListAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetRequest struct{}"
	}

	return strings.Join([]string{"ListAssetRequest", string(data)}, " ")
}
