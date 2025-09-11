package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListKmsTdeKeyResponseBodyKeyDetails struct {

	// **参数解释**: 秘钥ID。 **取值范围**: 不涉及。
	KeyId *string `json:"key_id,omitempty"`

	// **参数解释**: 默认主密钥标识。 **取值范围**: 默认主密钥标识为1，非默认标识为0。
	DefaultKeyFlag *string `json:"default_key_flag,omitempty"`

	// **参数解释**: 密钥别名。 **取值范围**: 不涉及。
	KeyAlias *string `json:"key_alias,omitempty"`

	// **参数解释**: 密钥生成算法。 **取值范围**: - AES_256 - SM4 - RSA_2048 - RSA_3072 - RSA_4096 - EC_P256 - EC_P384 - SM2 - ALL
	KeySpec *string `json:"key_spec,omitempty"`

	// **参数解释**: 用户域ID。 **取值范围**: 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释**: 秘钥状态。 **取值范围**: - “1”表示待激活状态。 - “2”表示启用状态。 - “3”表示禁用状态。 - “4”表示计划删除状态。 - “5”表示等待导入状态。
	KeyState *string `json:"key_state,omitempty"`
}

func (o ListKmsTdeKeyResponseBodyKeyDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsTdeKeyResponseBodyKeyDetails struct{}"
	}

	return strings.Join([]string{"ListKmsTdeKeyResponseBodyKeyDetails", string(data)}, " ")
}
