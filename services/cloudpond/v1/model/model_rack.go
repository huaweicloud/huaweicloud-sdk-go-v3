package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Rack 机柜
type Rack struct {

	// 机柜ID
	Id *string `json:"id,omitempty"`

	// 机柜名称
	Name *string `json:"name,omitempty"`

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	// 机柜目录ID
	RackCategoryId *string `json:"rack_category_id,omitempty"`

	// 机柜类型
	RackType *string `json:"rack_type,omitempty"`

	Status *RackStatus `json:"status,omitempty"`

	// 已分配存储容量
	StorageAssignedSize *int32 `json:"storage_assigned_size,omitempty"`

	// 机柜描述
	Description *string `json:"description,omitempty"`

	// 机柜SN号
	RackSnNo *string `json:"rack_sn_no,omitempty"`

	// 机柜位置号
	RackLocationNo *string `json:"rack_location_no,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 生效时间
	EffectedAt *sdktime.SdkTime `json:"effected_at,omitempty"`

	MarketOptions *MarketOptions `json:"market_options,omitempty"`

	// 计算单元信息
	ComputeUnit *[]ComputeSpec `json:"compute_unit,omitempty"`

	StorageUnit *StorageUnit `json:"storage_unit,omitempty"`

	RackInfo *RackInfo `json:"rack_info,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`
}

func (o Rack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rack struct{}"
	}

	return strings.Join([]string{"Rack", string(data)}, " ")
}
