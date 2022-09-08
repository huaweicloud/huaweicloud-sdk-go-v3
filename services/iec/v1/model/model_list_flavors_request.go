package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {

	// 页码。 当前页面数，默认为1。 取值大于等于0，取值为0时返回第1页。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回边缘实例规格列表当前页面的数量 。 取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件，规格的名称。
	Name *string `json:"name,omitempty"`

	// 边缘规格所在大区。  大小写通用，皆支持。 支持多个查询，中间使用“,”分隔。
	Area *string `json:"area,omitempty"`

	// 边缘规格所在省份。  大小写通用，皆支持。 支持多个查询，中间使用“,”分隔。
	Province *string `json:"province,omitempty"`

	// 边缘规格所在城市。  大小写通用，皆支持。 支持多个查询，中间使用“,”分隔。
	City *string `json:"city,omitempty"`

	// 边缘规格支持运营商。  大小写通用，皆支持。 支持多个查询，中间使用“,”分隔。
	Operator *string `json:"operator,omitempty"`

	// 查询条件，规格的ID。
	Id *string `json:"id,omitempty"`

	// 查询条件，边缘规格站点列表，站点之间用“,”分隔。
	SiteIds *string `json:"site_ids,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
