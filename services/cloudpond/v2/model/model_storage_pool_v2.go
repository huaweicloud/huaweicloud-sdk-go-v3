package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StoragePoolV2 存储池
type StoragePoolV2 struct {

	// 存储池ID
	Id *string `json:"id,omitempty"`

	// 存储池名称
	Name *string `json:"name,omitempty"`

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	// 存储类型
	StorageType *string `json:"storage_type,omitempty"`

	Status *StoragePoolStatus `json:"status,omitempty"`

	// 存储池大小。 当前已购买的存储容量。
	AssignedSize *int32 `json:"assigned_size,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	// 总容量
	Capacity *int32 `json:"capacity,omitempty"`

	MarketOptions *MarketOptions `json:"market_options,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 扩容场景预生成存储池关联的存储池ID
	ParentId *string `json:"parent_id,omitempty"`
}

func (o StoragePoolV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoragePoolV2 struct{}"
	}

	return strings.Join([]string{"StoragePoolV2", string(data)}, " ")
}
