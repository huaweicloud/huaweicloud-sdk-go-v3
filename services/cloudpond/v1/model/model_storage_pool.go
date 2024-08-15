package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StoragePool 存储池
type StoragePool struct {

	// 机柜ID
	Id *string `json:"id,omitempty"`

	// 存储池名称
	Name *string `json:"name,omitempty"`

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	StorageType *StorageType `json:"storage_type,omitempty"`

	Status *StoragePoolStatus `json:"status,omitempty"`

	// 存储池大小。 当前购买的存储容量。
	AssignedSize *int32 `json:"assigned_size,omitempty"`

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	// 总容量
	Capacity *int32 `json:"capacity,omitempty"`

	MarketOptions *MarketOptions `json:"market_options,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 生效时间
	EffectedAt *sdktime.SdkTime `json:"effected_at,omitempty"`
}

func (o StoragePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoragePool struct{}"
	}

	return strings.Join([]string{"StoragePool", string(data)}, " ")
}
