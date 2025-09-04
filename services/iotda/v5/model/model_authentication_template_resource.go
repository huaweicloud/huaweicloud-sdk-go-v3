package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthenticationTemplateResource 鉴权模板设备资源详情结构体。
type AuthenticationTemplateResource struct {

	// 设备ID，json对象只能使用FunctionDefinition下的函数从parameters中获取设备ID的取值。
	DeviceId *interface{} `json:"device_id"`

	Timestamp *TimestampResource `json:"timestamp,omitempty"`

	// mqtt认证密码，该字段只在设备认证方式为密码认证时生效，证书认证时无效，只能使用FunctionDefinition下的函数从parameters中编程密码的生成方式，平台比较函数解析结果与设备mqtt连接时携带的password是否相等，相等则认证通过。函数中必须包含设备原始秘钥参数${iotda::device::secret}，且只能在hash函数中使用。
	Password *interface{} `json:"password,omitempty"`
}

func (o AuthenticationTemplateResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthenticationTemplateResource struct{}"
	}

	return strings.Join([]string{"AuthenticationTemplateResource", string(data)}, " ")
}
