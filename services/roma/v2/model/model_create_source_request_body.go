package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSourceRequestBody struct {
	// 产品ID，自动向下取整

	ProductId int32 `json:"product_id"`
	// 设备ID，自动向下取整，不填为全部设备

	DeviceId *int32 `json:"device_id,omitempty"`
	// 主题，当设备ID为空时为产品级主题，设备ID不为空时为设备级主题

	Topic string `json:"topic"`
	// 是否payload使用base64，0-是 1-否

	IsBase64 *int32 `json:"is_base64,omitempty"`
	// 是否包含设备信息是否包含设备信息，0-是 1-否

	ContainDeviceInfo *int32 `json:"contain_device_info,omitempty"`
}

func (o CreateSourceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSourceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSourceRequestBody", string(data)}, " ")
}
