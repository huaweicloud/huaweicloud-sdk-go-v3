package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HarmonySoftBusResponseDto 鸿蒙软总线信息结构体，创建、查询鸿蒙软总线时返回
type HarmonySoftBusResponseDto struct {

	// 鸿蒙软总线ID，用于唯一标识一个鸿蒙软总线，在创建鸿蒙软总线时由物联网平台分配。
	BusId *string `json:"bus_id,omitempty"`

	// 鸿蒙软总线名称，单个资源空间下不可重复。
	BusName *string `json:"bus_name,omitempty"`

	// 设备组ID。
	GroupId *string `json:"group_id,omitempty"`

	// **参数说明**：资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的设备组归属到哪个资源空间下，否则创建的设备组将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	AppId *string `json:"app_id,omitempty"`

	// 鸿蒙软总线组信息同步状态，结果范围：Synchronized，Synchronizing，ToBeSynchronized - Synchronized：已同步。 - Synchronizing：同步中。 - ToBeSynchronized：待同步。
	Status *string `json:"status,omitempty"`
}

func (o HarmonySoftBusResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HarmonySoftBusResponseDto struct{}"
	}

	return strings.Join([]string{"HarmonySoftBusResponseDto", string(data)}, " ")
}
