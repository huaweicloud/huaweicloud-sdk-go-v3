package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneModel 云手机规格信息。
type PhoneModel struct {

	// 云手机服务器的规格名称，不超过64个字节
	ServerModelName *string `json:"server_model_name,omitempty"`

	// 云手机的规格名称，不超过64个字节
	PhoneModelName *string `json:"phone_model_name,omitempty"`

	// 规格状态 - 1 表示正常使用状态 - 0 表示已下线状态 已下线的规格不可用来购买云手机服务器
	Status *int32 `json:"status,omitempty"`

	// CPU核数
	Cpu *int32 `json:"cpu,omitempty"`

	// 内存大小，单位：MB
	Memory *int32 `json:"memory,omitempty"`

	// 系统存储大小，单位：GB
	Disk *int32 `json:"disk,omitempty"`

	// 分辨率，不超过16个字节
	Resolution *string `json:"resolution,omitempty"`

	// 扩展描述，不超过512个字节
	ExtendSpec *string `json:"extend_spec,omitempty"`

	// 规格名称，不超过64个字节
	SpecCode *string `json:"spec_code,omitempty"`

	// 当前云手机规格包含的云手机个数
	PhoneCapacity *int32 `json:"phone_capacity,omitempty"`

	// 镜像类型，只支持如下类型： - qemu_phone - cloud_phone - cloud_game
	ImageLabel *string `json:"image_label,omitempty"`

	// 产品类型 - 0：云手机 - 1：云手游
	ProductType *int32 `json:"product_type,omitempty"`
}

func (o PhoneModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneModel struct{}"
	}

	return strings.Join([]string{"PhoneModel", string(data)}, " ")
}
