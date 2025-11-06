package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Drive 驱动器详细信息，如型号、厂商等
type Drive struct {

	// 驱动器资源的ID
	Id *string `json:"id,omitempty"`

	// 驱动器资源的名称
	Name *string `json:"name,omitempty"`

	// 驱动器的制造商
	Manufacturer *string `json:"manufacturer,omitempty"`

	// 驱动器型号
	Model *string `json:"model,omitempty"`

	// 驱动器的序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 驱动器的版本信息
	Revision *string `json:"revision,omitempty"`

	// 驱动器的介质类型
	MediaType *string `json:"media_type,omitempty"`

	// 驱动器遵从的协议
	Protocol *string `json:"protocol,omitempty"`

	// SAS地址
	SasAddress *string `json:"sas_address,omitempty"`

	// 容量（单位：byte）
	CapacityBytes *int64 `json:"capacity_bytes,omitempty"`

	// 驱动器接口的最大速率（单位：Gbps）
	CapableSpeedGbs *int32 `json:"capable_speed_gbs,omitempty"`

	// 驱动器接口的协商速率（单位：Gbps）
	NegotiatedSpeedGbs *int32 `json:"negotiated_speed_gbs,omitempty"`

	// 驱动器的热备状态
	HotspareType *string `json:"hotspare_type,omitempty"`

	// 驱动器上电运行时间（单位：h）
	HoursOfPoweredUp *int32 `json:"hours_of_powered_up,omitempty"`

	// 驱动器电源状态
	PowerState *string `json:"power_state,omitempty"`

	// 驱动器巡检状态
	PatrolState *string `json:"patrol_state,omitempty"`

	// 驱动器的资源归属
	AssociatedResource *string `json:"associated_resource,omitempty"`

	// 驱动器的剩余寿命百分比
	PredictedMediaLifeLeftPercent *int32 `json:"predicted_media_life_left_percent,omitempty"`

	// 健康状态
	Health *string `json:"health,omitempty"`
}

func (o Drive) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Drive struct{}"
	}

	return strings.Join([]string{"Drive", string(data)}, " ")
}
