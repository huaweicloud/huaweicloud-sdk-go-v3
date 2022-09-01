package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DevicesCalculation struct {
	NormalDevices *Calculation `json:"normal_devices,omitempty" xml:"normal_devices"`

	GatewayDevices *Calculation `json:"gateway_devices,omitempty" xml:"gateway_devices"`

	SubsetsDevices *Calculation `json:"subsets_devices,omitempty" xml:"subsets_devices"`
}

func (o DevicesCalculation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevicesCalculation struct{}"
	}

	return strings.Join([]string{"DevicesCalculation", string(data)}, " ")
}
