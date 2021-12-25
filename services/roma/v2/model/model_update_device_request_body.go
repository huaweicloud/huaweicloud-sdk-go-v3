package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDeviceRequestBody struct {
	// 设备名称，支持中文、中文标点符号（）。；，：“”、？《》及英文大小写、数字及英文符号()_,#.?'-@%&!, 长度2-64

	DeviceName string `json:"device_name"`
	// 设备状态 0启用 1禁用

	Status int32 `json:"status"`
	// 备注

	Description *string `json:"description,omitempty"`
	// 标签

	Tags *[]string `json:"tags,omitempty"`
}

func (o UpdateDeviceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDeviceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDeviceRequestBody", string(data)}, " ")
}
