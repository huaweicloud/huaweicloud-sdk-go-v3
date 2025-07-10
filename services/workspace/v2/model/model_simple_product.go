package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleProduct 产品信息。
type SimpleProduct struct {

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 产品类型。  - BASE：表示产品基础套餐，套餐镜像中不包括除操作系统之外的其他商业软件，私有镜像场景只能使用此类套餐。
	Type *string `json:"type,omitempty"`

	// CPU。
	Cpu *string `json:"cpu,omitempty"`

	// 内存。
	Memory *string `json:"memory,omitempty"`

	// 产品描述。
	Descriptions *string `json:"descriptions,omitempty"`

	// 周期套餐标识，0表示包周期，1表示按需。
	ChargeMode *string `json:"charge_mode,omitempty"`
}

func (o SimpleProduct) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleProduct struct{}"
	}

	return strings.Join([]string{"SimpleProduct", string(data)}, " ")
}
