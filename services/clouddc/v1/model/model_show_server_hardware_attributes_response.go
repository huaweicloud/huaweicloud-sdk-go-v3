package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerHardwareAttributesResponse Response Object
type ShowServerHardwareAttributesResponse struct {
	Summary *HardwareSummary `json:"summary,omitempty"`

	// 内存列表
	Memorys *[]Memory `json:"memorys,omitempty"`

	// cpu列表
	Processors *[]Processors `json:"processors,omitempty"`

	// 网络适配器详细信息
	NetworkAdapters *[]NetworkAdapter `json:"network_adapters,omitempty"`

	// 风扇列表
	Fans *[]Fan `json:"fans,omitempty"`

	// 电源列表
	Powers *[]Power `json:"powers,omitempty"`

	// 存储控制器列表
	StorageControllers *[]StorageController `json:"storage_controllers,omitempty"`
	HttpStatusCode     int                  `json:"-"`
}

func (o ShowServerHardwareAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerHardwareAttributesResponse struct{}"
	}

	return strings.Join([]string{"ShowServerHardwareAttributesResponse", string(data)}, " ")
}
