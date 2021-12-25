package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Source struct {
	// 源数据源ID

	SourceId *int32 `json:"source_id,omitempty"`
	// 产品ID

	ProductId *int32 `json:"product_id,omitempty"`
	// 设备ID，不填为全部设备

	DeviceId *int32 `json:"device_id,omitempty"`
	// 主题，当设备ID为空时为产品级主题，设备ID不为空时为设备级主题

	Topic *string `json:"topic,omitempty"`
	// 设备名称

	DeviceName *string `json:"device_name,omitempty"`
	// 产品名称

	ProductName *string `json:"product_name,omitempty"`
	// 是否payload使用base64，0-是 1-否

	IsBase64 *int32 `json:"is_base64,omitempty"`
	// 是否包含设备信息，0-是，1-否

	ContainDeviceInfo *int32 `json:"contain_device_info,omitempty"`
}

func (o Source) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Source struct{}"
	}

	return strings.Join([]string{"Source", string(data)}, " ")
}
