package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘hub与外部实体的mqtt连接信息
type MqttConnectionInfo struct {

	// 采用cleint方式连接时，mqtt服务器地址
	ServerAddress *string `json:"server_address,omitempty" xml:"server_address"`

	// mqtt连接时，client_id
	ClientId *string `json:"client_id,omitempty" xml:"client_id"`

	// 鉴权类型。支持密钥认证接入(SECRET)和证书认证接入(CERTIFICATES)两种方式。使用密钥认证接入方式(SECRET)填写user_name和user_name字段，使用证书认证接入方式(CERTIFICATES)填写privateKey和certificate字段
	AuthType *string `json:"auth_type,omitempty" xml:"auth_type"`

	// 证书秘钥
	PrivateKey *string `json:"private_key,omitempty" xml:"private_key"`

	// 证书
	Certificate *string `json:"certificate,omitempty" xml:"certificate"`

	// 用户名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 密码
	Password *string `json:"password,omitempty" xml:"password"`

	// 服务质量,默认为0,表示最多一次的传输,1表示至少一次,2表示仅一次.
	Qos *int32 `json:"qos,omitempty" xml:"qos"`
}

func (o MqttConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MqttConnectionInfo struct{}"
	}

	return strings.Join([]string{"MqttConnectionInfo", string(data)}, " ")
}
