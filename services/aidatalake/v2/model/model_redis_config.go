package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisConfig **参数解释**：Redis集群配置，目前仅支持DCS。
type RedisConfig struct {

	// **参数解释**：账号名称。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字及特殊字符，长度为1~64个字符。 **默认取值**：default。
	Username *string `json:"username,omitempty"`

	// **参数解释**：DEW的凭据名称。存储着DCS账号的密码，格式为：凭据键为password，值为实际的密码。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字及特殊字符，长度为1~64个字符。 **默认取值**：不涉及。
	DewSecretName string `json:"dew_secret_name"`

	// **参数解释**：是否开启SSL。 **约束限制**：不涉及。 **取值范围**：   - true：开启SSL。   - false：关闭SSL。 **默认取值**：true。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数解释**：DCS实例的ID。获取方式：在DCS控制台实例详情页查看，或通过调用查询DCS实例列表接口获取。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字及中划线，长度为4~1024个字符。 **默认取值**：不涉及。
	DcsId string `json:"dcs_id"`

	// **参数解释**：VPCEP服务端ID，用于连接客户侧网络。获取方式：在VPCEP控制台查看，或通过调用查询VPCEP服务列表接口获取。开启GCS高可用时必填，用于通过VPCEP方式连接客户侧DCS。 **约束限制**：开启GCS高可用时必填。 **取值范围**：不涉及。 **默认取值**：不涉及。
	VpcepServiceId string `json:"vpcep_service_id"`
}

func (o RedisConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisConfig struct{}"
	}

	return strings.Join([]string{"RedisConfig", string(data)}, " ")
}
