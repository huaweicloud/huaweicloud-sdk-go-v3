package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PropertyToUpdate 云手机属性信息。
type PropertyToUpdate struct {

	// 云手机id，不超过32个字节。
	PhoneId string `json:"phone_id"`

	// 云手机属性列表，为Json格式字符串。
	Property string `json:"property"`

	// 用户自定义属性键值对。若涉及 OS 系统属性，需遵循系统属性规范。注意：本字段与 property 字段的合并总长度不得超过 7800 字节。
	CustomProperty map[string]string `json:"custom_property,omitempty"`
}

func (o PropertyToUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyToUpdate struct{}"
	}

	return strings.Join([]string{"PropertyToUpdate", string(data)}, " ")
}
