package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StorageGearV2 StorageGear接口返回层对象模型
type StorageGearV2 struct {

	// 存储档位ID
	Id *string `json:"id,omitempty"`

	// 存储阶梯值, 如：35TB
	Gear *int32 `json:"gear,omitempty"`

	// 存储类型. SAS:高IO,SSD:超高IO [VS_SMALL_CAP,VS_MEDIUM_CAP,VS_LARGE_CAP 视图存储](tag:cmcc)
	StorageType *string `json:"storage_type,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	// 部署地区
	ZoneCode *string `json:"zone_code,omitempty"`

	SaleCycles *[]SaleCycle `json:"sale_cycles,omitempty"`
}

func (o StorageGearV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageGearV2 struct{}"
	}

	return strings.Join([]string{"StorageGearV2", string(data)}, " ")
}
