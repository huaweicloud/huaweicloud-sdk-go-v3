package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductRequest Request Object
type ListProductRequest struct {

	// 如果不为空，将按产品ID进行过滤后返回。
	ProductId *string `json:"product_id,omitempty"`

	// 如果不为空，将按规格ID进行过滤后返回。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 可用分区，如果不为空，将按可用分区进行过滤后返回。 - 获取方式详见可用区管理ListAvailabilityZone：\"GET  /v1/{project_id}/availability-zone\"。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 产品套餐的操作系统类型，当前支持：Windows。 - Linux - Windows - Other
	OsType *string `json:"os_type,omitempty"`

	// 套餐标识。 - 1：表示包周期。 - 0：表示按需。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 架构类型，当前支持：x86。 - x86 - arm
	Architecture *string `json:"architecture,omitempty"`

	// 套餐类型： - general：表示产品通用套餐。 - dedicated：表示产品专属主机套餐。
	PackageType *string `json:"package_type,omitempty"`
}

func (o ListProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductRequest struct{}"
	}

	return strings.Join([]string{"ListProductRequest", string(data)}, " ")
}
