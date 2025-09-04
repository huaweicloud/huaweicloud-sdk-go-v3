package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddHarmonySoftBusDto 创建鸿蒙软总线组请求结构体
type AddHarmonySoftBusDto struct {

	// **参数说明**：鸿蒙软总线组名称，单个资源空间下不可重复。 **取值范围**：长度不超过64，只允许中文、字母、数字、以及_? '#().,&%@!-等字符的组合。
	BusName *string `json:"bus_name,omitempty"`

	// **参数说明**：设备组ID，用于唯一标识一个设备组，在创建设备组时由物联网平台分配。
	GroupId *string `json:"group_id,omitempty"`

	// **参数说明**：资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的设备组归属到哪个资源空间下，否则创建的设备组将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。 **取值范围**：长度不超过36，只允许字母、数字、下划线（_）、连接符（-）的组合。
	AppId *string `json:"app_id,omitempty"`
}

func (o AddHarmonySoftBusDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHarmonySoftBusDto struct{}"
	}

	return strings.Join([]string{"AddHarmonySoftBusDto", string(data)}, " ")
}
