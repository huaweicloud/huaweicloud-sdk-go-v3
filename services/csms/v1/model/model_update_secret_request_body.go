package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecretRequestBody 更新凭据对象的元数据信息请求体
type UpdateSecretRequestBody struct {

	// 用于加密保护凭据值的KMS主密钥ID。更新凭据的主密钥后，仅新创建的凭据版本使用更新后的主密钥ID加密，之前的凭据版本依旧使用之前的主密钥ID解密。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	// 凭据的描述信息。 约束：2048字节。
	Description *string `json:"description,omitempty"`

	// 自动轮转  取值：true 开启 false 关
	AutoRotation *bool `json:"auto_rotation,omitempty"`

	// 轮转周期  约束：6小时-8,760小时 （365天）  类型：Integer[unit] ，Integer表示时间长度 。unit表示时间单位，d（天）、h（小时）、m（分钟）、s（秒）。例如 1d 表示一天，24h也表示一天  说明：当开启自动轮转时，必须填写该值
	RotationPeriod *string `json:"rotation_period,omitempty"`

	// 凭据订阅的事件列表，当前最大可订阅一个事件。当事件包含的基础事件触发时，通知消息将发送到事件对应的通知主题。
	EventSubscriptions *[]string `json:"event_subscriptions,omitempty"`

	// FunctionGraph函数的urn。
	RotationFuncUrn *string `json:"rotation_func_urn,omitempty"`
}

func (o UpdateSecretRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecretRequestBody", string(data)}, " ")
}
