package model

import (
	"encoding/json"

	"strings"
)

// 创建设备组请求结构体
type AddDeviceGroupDto struct {
	// 设备组名称，单个资源空间下不可重复。

	Name *string `json:"name,omitempty"`
	// 设备组描述。

	Description *string `json:"description,omitempty"`
	// 父设备组ID，携带该参数时表示在该设备组下创建一个子设备组。

	SuperGroupId *string `json:"super_group_id,omitempty"`
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，建议携带该参数指定创建的设备组归属到哪个资源空间下，否则创建的设备组将会归属到[默认资源空间](https://support.huaweicloud.com/usermanual-iothub/iot_01_0006.html#section0)下。

	AppId *string `json:"app_id,omitempty"`
}

func (o AddDeviceGroupDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddDeviceGroupDto struct{}"
	}

	return strings.Join([]string{"AddDeviceGroupDto", string(data)}, " ")
}
