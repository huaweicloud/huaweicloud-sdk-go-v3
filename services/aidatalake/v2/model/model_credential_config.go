package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CredentialConfig **参数解释**：Endpoint级别用户委托名称，用于Endpoint的存储等配置权限认证，以及作为作业运行时的默认委托。
type CredentialConfig struct {

	// **参数解释**：Endpoint级别用户委托名称，用于Endpoint的存储等配置权限认证，以及作为作业运行时的默认委托。 **约束限制**：不涉及。 **取值范围**：支持英文字母、数字及特殊字符，长度为1~64个字符。 **默认取值**：不涉及。
	AgencyName *string `json:"agency_name,omitempty"`
}

func (o CredentialConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CredentialConfig struct{}"
	}

	return strings.Join([]string{"CredentialConfig", string(data)}, " ")
}
