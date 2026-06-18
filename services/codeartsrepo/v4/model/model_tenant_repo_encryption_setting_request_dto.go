package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TenantRepoEncryptionSettingRequestDto struct {

	// **参数解释：** 租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 加密类型。 **取值范围：** KMS表示开启KMS加密，normal或者null表示未开启KMS加密。
	EncryptionType *string `json:"encryption_type,omitempty"`

	// **参数解释：** 是否开启租户下默认加密设置。
	DefaultEncryptionEnabled *bool `json:"default_encryption_enabled,omitempty"`

	// **参数解释：** 加密主密钥的名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	CmkKeyName *string `json:"cmk_key_name,omitempty"`

	// **参数解释：** 加密主密钥的id。 **取值范围：** 字符串长度不少于1，不超过1000。
	CmkKeyId *string `json:"cmk_key_id,omitempty"`
}

func (o TenantRepoEncryptionSettingRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantRepoEncryptionSettingRequestDto struct{}"
	}

	return strings.Join([]string{"TenantRepoEncryptionSettingRequestDto", string(data)}, " ")
}
