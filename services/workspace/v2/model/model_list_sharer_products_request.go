package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharerProductsRequest Request Object
type ListSharerProductsRequest struct {

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 协同方数。该套餐支持的最大协同人数。
	ShareSpaceSize *string `json:"share_space_size,omitempty"`

	// 周期套餐标识。0表示包周期，1表示按需, 6表示一次性计费。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 是否是GPU套餐。1表示GPU套餐，0表示非GPU套餐，默认null查询所有类型。
	IsGpu *int32 `json:"is_gpu,omitempty"`

	// 套餐系列。user_sharer表示用户协同套餐，desktop_sharer表示桌面协同套餐。
	PackageType *string `json:"package_type,omitempty"`

	// 每页数量，范围0-100，默认100。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSharerProductsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharerProductsRequest struct{}"
	}

	return strings.Join([]string{"ListSharerProductsRequest", string(data)}, " ")
}
