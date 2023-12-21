package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDetailAzInfo 可用区信息。
type InstanceDetailAzInfo struct {

	// 云堡垒机实例所在可用区ID。
	Region string `json:"region"`

	// 云堡垒机实例所在可用分区ID。(实例为主备模式时作为主机实例所在可用分区)
	Zone string `json:"zone"`

	// 云堡垒机实例所在可用分区中文名称。(实例为主备模式时作为主机实例所在可用分区中文名称)
	AvailabilityZoneDisplay string `json:"availability_zone_display"`

	// 云堡垒机备机实例所在可用区。
	SlaveZone *string `json:"slave_zone,omitempty"`

	// 云堡垒机备机实例所在可用区中文名称。
	SlaveZoneDisplay *string `json:"slave_zone_display,omitempty"`
}

func (o InstanceDetailAzInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetailAzInfo struct{}"
	}

	return strings.Join([]string{"InstanceDetailAzInfo", string(data)}, " ")
}
