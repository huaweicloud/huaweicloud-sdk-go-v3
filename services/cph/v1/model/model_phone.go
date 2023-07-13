package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Phone 云手机信息。
type Phone struct {

	// 云手机的名称，不超过65个字符
	PhoneName *string `json:"phone_name,omitempty"`

	// 云手机所在的服务器ID，不超过32个字节
	ServerId *string `json:"server_id,omitempty"`

	// 云手机的唯一标识，不超过32个字节
	PhoneId *string `json:"phone_id,omitempty"`

	// 云手机规格名称，不超过64个字节
	PhoneModelName *string `json:"phone_model_name,omitempty"`

	// 云手机镜像ID，不超过32个字节
	ImageId *string `json:"image_id,omitempty"`

	// 云手机是否开启VNC服务 - true表示开启 - false表示不开启
	VncEnable *string `json:"vnc_enable,omitempty"`

	// 云手机状态 - 0: 创建中 - 1：创建中 - 2：运行中 - 3：重置中 - 4：重启中 - 6：冻结 - 7：正在关机 - 8：已关机 - -5：重置失败 - -6：重启失败 - -7：手机异常 - -8：创建失败 - -9：关机失败
	Status *int32 `json:"status,omitempty"`

	// 云手机类型 - 0：普通云手机 - 1：试玩云手机
	Type *int32 `json:"type,omitempty"`

	Metadata *PhoneMetadata `json:"metadata,omitempty"`

	// 创建时间 时间格式为UTC
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间 时间格式为UTC
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o Phone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Phone struct{}"
	}

	return strings.Join([]string{"Phone", string(data)}, " ")
}
