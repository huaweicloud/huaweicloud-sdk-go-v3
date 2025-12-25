package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcsHypervisor 虚拟化信息。
type HwcEcsHypervisor struct {

	// 虚拟化类型。
	HypervisorType *string `json:"hypervisor_type,omitempty"`

	// 预留属性。
	CsdHypervisor *string `json:"csd_hypervisor,omitempty"`
}

func (o HwcEcsHypervisor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcsHypervisor struct{}"
	}

	return strings.Join([]string{"HwcEcsHypervisor", string(data)}, " ")
}
