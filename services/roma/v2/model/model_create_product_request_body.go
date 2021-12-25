package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProductRequestBody struct {
	// 应用ID

	AppId string `json:"app_id"`
	// 产品名称，创建产品时租户内唯一，长度最大64，仅支持中文，英文字母，数字，下划线和中划线

	Name string `json:"name"`
	// 产品供应商ID

	ManufacturerId string `json:"manufacturer_id"`
	// 厂商名称

	ManufacturerName string `json:"manufacturer_name"`
	// 产品型号

	Model string `json:"model"`
	// 产品类型，0-普通产品(不支持子设备) 1-网关产品

	ProductType int32 `json:"product_type"`
	// 产品描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 产品的协议类型 0-mqtt 2-modbus 4-opcua

	ProtocolType int32 `json:"protocol_type"`
	// 产品的设备类型（默认Default）

	DeviceType *string `json:"device_type,omitempty"`
	// 关联产品模板ID（使用产品模板创建产品时使用，否则为空），自动向下取整

	TemplateId *int32 `json:"template_id,omitempty"`
	// 模型版本

	Version *string `json:"version,omitempty"`
	// 产品的数据格式 0-JSON 1-USER_DEFINED

	DataFormat *int32 `json:"data_format,omitempty"`
}

func (o CreateProductRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProductRequestBody", string(data)}, " ")
}
