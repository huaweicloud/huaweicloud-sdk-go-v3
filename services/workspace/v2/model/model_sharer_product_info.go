package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SharerProductInfo struct {

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 是否是GPU类型的规格。
	IsGpu *bool `json:"is_gpu,omitempty"`

	// 产品描述。
	Descriptions *string `json:"descriptions,omitempty"`

	// 周期套餐标识。0表示包周期，1表示按需, 6表示一次性计费。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 资源规格。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务编码。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 套餐类型。 - user_sharer：用户协同套餐 - desktop_sharer: 桌面协同套餐
	PackageType *string `json:"package_type,omitempty"`

	// 产品名称<语言，各语言对应的产品名>。
	Name map[string]string `json:"name,omitempty"`

	// 协同方数。该套餐支持的最大协同人数。
	ShareSpaceSize *int32 `json:"share_space_size,omitempty"`
}

func (o SharerProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SharerProductInfo struct{}"
	}

	return strings.Join([]string{"SharerProductInfo", string(data)}, " ")
}
