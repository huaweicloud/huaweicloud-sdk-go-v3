package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateProductRequestBody struct {

	// 产品名称，创建产品时租户内唯一，长度1-64，仅支持中文，英文字母，数字，下划线和中划线
	Name string `json:"name" xml:"name"`

	// 产品描述，长度0-200
	Description *string `json:"description,omitempty" xml:"description"`

	// 产品供应商ID，支持英文大小写，数字，下划线和中划线，长度2-50
	ManufacturerId string `json:"manufacturer_id" xml:"manufacturer_id"`

	// 厂商名称，支持长度2-64
	ManufacturerName string `json:"manufacturer_name" xml:"manufacturer_name"`

	// 产品型号，支持英文大小写，数字，下划线，中划线和空格(首尾空格会被忽略)，长度2-50
	Model string `json:"model" xml:"model"`

	// 产品的设备类型（默认Default Type）
	DeviceType string `json:"device_type" xml:"device_type"`

	// 模型版本
	Version *string `json:"version,omitempty" xml:"version"`
}

func (o UpdateProductRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProductRequestBody", string(data)}, " ")
}
