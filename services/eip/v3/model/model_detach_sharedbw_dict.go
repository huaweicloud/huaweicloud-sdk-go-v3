package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachSharedbwDict 弹性公网IP移出共享带宽后带宽的参数
type DetachSharedbwDict struct {

	// - 功能说明：带宽名称
	Name *string `json:"name,omitempty"`

	// - 功能说明：带宽大小
	Size int32 `json:"size"`

	// - 功能说明：带宽计费模式
	ChargeMode string `json:"charge_mode"`
}

func (o DetachSharedbwDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachSharedbwDict struct{}"
	}

	return strings.Join([]string{"DetachSharedbwDict", string(data)}, " ")
}
