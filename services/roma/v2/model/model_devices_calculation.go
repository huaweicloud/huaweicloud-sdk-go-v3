package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DevicesCalculation struct {
	NormalDevices *Calculation `json:"normal_devices,omitempty"`

	GatewayDevices *Calculation `json:"gateway_devices,omitempty"`

	SubsetsDevices *Calculation `json:"subsets_devices,omitempty"`
}

func (o DevicesCalculation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicesCalculation struct{}"
	}

	return strings.Join([]string{"DevicesCalculation", string(data)}, " ")
}
