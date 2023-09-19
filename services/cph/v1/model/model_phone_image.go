package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PhoneImage 云手机镜像信息。
type PhoneImage struct {

	// 手机镜像名称，不超过128个字节。
	ImageName *string `json:"image_name,omitempty"`

	// 镜像操作系统类型，不超过16个字节。
	OsType *string `json:"os_type,omitempty"`

	// 镜像类型。 - 1：公有镜像 - 2 ：私有镜像
	IsPublic *int32 `json:"is_public,omitempty"`

	// 手机操作系统，不超过36个字节。
	OsName *string `json:"os_name,omitempty"`

	// 镜像适用的云手机规格。 - cloud_phone：适用于physical.rx1.xlarge 类型云手机服务器 - cloud_phone_1620：适用于physical.kg1.4xlarge.cp类型云手机服务器 - cloud_game：适用于physical.rx1.xlarge.cg 类型云手游服务器 - cloud_game_1620：适用于physical.kg1.4xlarge.cg 类型云手游服务器 - qemu_phone： 适用于physical.rx1.xlarge 类型云手机服务器中 qemu类型云手机规格
	ImageLabel *string `json:"image_label,omitempty"`

	// 手机镜像唯一标识ID，不超过32个字节。
	ImageId *string `json:"image_id,omitempty"`
}

func (o PhoneImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhoneImage struct{}"
	}

	return strings.Join([]string{"PhoneImage", string(data)}, " ")
}
